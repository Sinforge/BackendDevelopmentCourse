package repository

import (
	"context"
	"io"
	"mime/multipart"
	"net/textproto"
	"pr9/internal/models"
	"pr9/pkg/mongoconnector"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"gopkg.in/mgo.v2"
)

type MongoRepo struct {
	db      *mongoconnector.Connector
	grdConn *mgo.GridFS
}

func New(conn *mongoconnector.Connector, gridFs *mgo.GridFS) Repository {
	return &MongoRepo{
		db:      conn,
		grdConn: gridFs,
	}
}

func (m MongoRepo) SaveFileDataToGridFS(ctx context.Context, file *multipart.FileHeader) error {
	uploadStream, err := m.grdConn.Create(file.Filename)
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
	session := m.db.Session.Copy()
	defer session.Close()

	fileID, err := primitive.ObjectIDFromHex(file.ID)
	if err != nil {
		return err
	}

	err = m.grdConn.Files.UpdateId(fileID, bson.M{
		"$set": bson.M{
			"metadata": bson.M{
				"original_filename": file.Data.Filename,
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (m MongoRepo) GetFilesList(ctx context.Context) ([]models.File, error) {
	session := m.db.Session.Copy()
	defer session.Close()

	var files []models.File

	iter := m.grdConn.Find(nil).Iter()
	var result gridfs.File
	for iter.Next(&result) {
		log.Info(string(result.Metadata))
		files = append(files, models.File{
			ID: result.ID.(primitive.ObjectID).Hex(),
			Data: &multipart.FileHeader{
				Filename: result.Name,
				Header:   textproto.MIMEHeader{},
				Size:     int64(result.ChunkSize),
			},
		})
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return files, nil
}

func (m MongoRepo) UpdateFileByID(ctx context.Context, file models.File) error {
	session := m.db.Session.Copy()
	defer session.Close()

	fileID, err := primitive.ObjectIDFromHex(file.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": fileID}
	update := bson.M{"$set": bson.M{"name": "alohadance_2"}}

	err = m.grdConn.Files.Update(filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m MongoRepo) DeleteFileByID(ctx context.Context, fileID string) error {
	session := m.db.Session.Copy()
	defer session.Close()

	fileIDD, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}

	err = m.grdConn.RemoveId(fileIDD)
	if err != nil {
		return err
	}

	return nil
}

func (m MongoRepo) GetFileByID(ctx context.Context, fileID string) (file models.File, err error) {
	session := m.db.Session.Copy()
	defer session.Close()

	fileIDD, err := primitive.ObjectIDFromHex(file.ID)
	if err != nil {
		return file, err
	}

	var result gridfs.File
	err = m.grdConn.Find(bson.M{"_id": fileIDD}).One(&result)
	if err != nil {
		return file, err
	}

	file = models.File{
		ID: result.ID.(primitive.ObjectID).Hex(),
		Data: &multipart.FileHeader{
			Filename: result.Name,
			Header:   textproto.MIMEHeader{},
			Size:     int64(result.ChunkSize),
		},
	}

	return file, nil
}
