package repository

import (
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"pr9/internal/models"
	"pr9/pkg/mongoconnector"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	db      *mongo.Database
	grdConn *gridfs.Bucket
}

func New(conn *mongoconnector.Connector) Repository {
	db := conn.Database
	grdBucket, err := gridfs.NewBucket(
		db,
	)
	if err != nil {
		return nil
	}

	return &MongoRepo{
		db:      db,
		grdConn: grdBucket,
	}
}

func (m MongoRepo) SaveFileDataToGridFS(ctx context.Context, file *multipart.FileHeader) error {
	uploadStream, err := m.grdConn.OpenUploadStream(file.Filename)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	uploadedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	_, err = io.Copy(uploadStream, uploadedFile)
	if err != nil {
		return err
	}

	return nil
}

func (m MongoRepo) SaveFileInfo(ctx context.Context, file models.File) error {
	fileID, err := primitive.ObjectIDFromHex(file.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": fileID}
	update := bson.M{
		"$set": bson.M{
			"metadata": bson.M{
				"original_filename": file.Data,
			},
		},
	}

	_, err = m.db.Collection("fs.files").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m MongoRepo) GetFilesList(ctx context.Context) ([]models.File, error) {
	cursor, err := m.db.Collection("fs.files").Find(ctx, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var files []models.File

	for cursor.Next(ctx) {
		var result gridfs.File
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		files = append(files, models.File{
			ID:   result.ID.(primitive.ObjectID).Hex(),
			Name: result.Name,
			Data: result.Metadata,
		})
	}

	return files, nil
}

func (m MongoRepo) UpdateFileByID(ctx context.Context, file models.File) error {
	fileID, err := primitive.ObjectIDFromHex(file.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": fileID}
	update := bson.M{"$set": bson.M{"name": file.Name}}

	_, err = m.db.Collection("fs.files").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m MongoRepo) DeleteFileByID(ctx context.Context, fileID string) error {
	fileIDD, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}

	err = m.grdConn.Delete(fileIDD)
	if err != nil {
		return err
	}

	return nil
}

//func (m MongoRepo) GetFileByID(ctx context.Context, fileID string) (file models.File, err error) {
//	objID, err := primitive.ObjectIDFromHex("6547dac7e01c9e7744a44d39")
//	if err != nil {
//		return file, err
//	}
//
//	stream, err := m.grdConn.OpenDownloadStream(bson.M{"_id": objID})
//	if err != nil {
//		return file, err
//	}
//
//	data, err := ioutil.ReadAll(stream)
//	if err != nil {
//		return file, err
//	}
//
//	log.Println(data)
//
//	var result gridfs.File
//	cursor, err := m.grdConn.FindContext(ctx, bson.M{"_id": objID}, options.GridFSFind())
//	if err != nil {
//		return file, err
//	}
//
//	err = cursor.Decode(&result)
//	if err != nil {
//		return file, err
//	}
//
//	file = models.File{
//		ID: result.ID.(primitive.ObjectID).Hex(),
//		Data: &multipart.FileHeader{
//			Filename: result.Name,
//			Header:   textproto.MIMEHeader{},
//			Size:     int64(result.ChunkSize),
//		},
//	}
//
//	return file, nil
//}

func (m MongoRepo) GetFileByID(ctx context.Context, fileID string) (file models.File, err error) {
	// Convert the provided fileID string to an ObjectID
	objID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return file, err
	}

	// Try to open a download stream for the file with the specified ID
	stream, err := m.grdConn.OpenDownloadStream(objID)
	if err != nil {
		return file, err
	}

	// Read all data from the stream
	data, err := ioutil.ReadAll(stream)
	if err != nil {
		return file, err
	}

	return models.File{
		ID:   fileID,
		Data: data,
	}, nil
}
