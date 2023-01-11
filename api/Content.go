package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id      path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
/*
* Get List of Buckets from S3
 */
func (server *Server) GetBuckets(ctx *gin.Context) {

	result, err := server.s3.ListBuckets(nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Return the list of buckets in the response
	ctx.JSON(http.StatusOK, gin.H{"buckets from API Endpoint": result.Buckets})

}

/*
* Get List of contents in S3 of specific bucket
 */
func (server *Server) GetContentsByBucketName(ctx *gin.Context) {

	result, err := server.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(server.config.AWSS3BUCKETNAME),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Return the list of buckets in the response
	ctx.JSON(http.StatusOK, gin.H{"buckets from API Endpoint": result.Contents})

}

type UploadObj struct {
	FileObj    *multipart.FileHeader `form:"data" binding:"required"`
	Name       string                `form:"name" binding:"required"`
	Size       int                   `form:"size" binding:"required"`
	Created_At string                `form:"created_at" binding:"required"`
	FileId     string                `form:"fileid" binding:"required"`
}

func (server *Server) Upload(ctx *gin.Context) {
	mediaTypes := [3]string{"video/mp4", "video/mpg", "text/plain"}
	var req UploadObj
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	contentType := req.FileObj.Header.Get("Content-Type")

	matched := false
	for _, t := range mediaTypes {
		if t == contentType {
			matched = true
			break
		}
	}

	if !matched {
		ctx.JSON(http.StatusUnsupportedMediaType,
			"Unsupported Media Type")
		return
	}
	result, _ := server.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(server.config.AWSS3BUCKETNAME),
		Key:    aws.String(req.FileObj.Filename),
	})

	if result.Body != nil {
		ctx.JSON(http.StatusConflict,
			"File exists")
		return
	}

	file, err := req.FileObj.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	defer file.Close()
	// Read the contents of the file into a slice of bytes
	content, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.s3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(server.config.AWSS3BUCKETNAME),
		Key:         aws.String(req.FileObj.Filename),
		ContentType: aws.String(req.FileObj.Header.Get("Content-Type")),
		Body:        bytes.NewReader(content),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, gin.H{"The File Content": content})

}

type RequestGetContentById struct {
	Id string `uri:"id" binding:"required"`
}

/*
* Get Content by Id.
* Returns 404 if Id does not exist.
 */
func (server *Server) GetContentById(ctx *gin.Context) {

	var req RequestGetContentById
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Use the GetObject() function to retrieve the object
	result, err := server.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(server.config.AWSS3BUCKETNAME),
		Key:    aws.String(req.Id),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "File not found!")
		return
	}

	if err != nil {
		panic(err)
	}

	body := result.Body
	defer body.Close()

	// Create a buffer to store the object data
	buffer := make([]byte, *result.ContentLength)
	_, err = body.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the object data
	fmt.Println(string(buffer))
	ctx.JSON(http.StatusOK, gin.H{"The File Content": string(buffer)})

}

/*
* Delete Content by Id.
* Returns 404 if Id does not exist.
 */
func (server *Server) RemoveContentById(ctx *gin.Context) {
	var req RequestGetContentById
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check if the object exists
	_, err := server.s3.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(server.config.AWSS3BUCKETNAME),
		Key:    aws.String(req.Id),
	})
	if err != nil {
		// Return an error if the object does not exist
		ctx.JSON(http.StatusBadRequest, "File not found!")
		return
	}

	// Delete the object
	_, err = server.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(server.config.AWSS3BUCKETNAME),
		Key:    aws.String(req.Id),
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error when deleting file!")
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"The File Content has been deleted successfully": req.Id})
}
