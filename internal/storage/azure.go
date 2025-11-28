package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/joho/godotenv"
)

var (
	accountName   string
	accountKey    string
	containerName string
	credential    *azblob.SharedKeyCredential
)

// -------------------------------------------------------------
// Initialize Azure Blob Storage
// -------------------------------------------------------------
func InitAzureBlob() {
	godotenv.Load(".env")

	accountName = os.Getenv("AZURE_STORAGE_ACCOUNT")
	accountKey = os.Getenv("AZURE_STORAGE_KEY")
	containerName = os.Getenv("AZURE_STORAGE_CONTAINER")

	var err error
	credential, err = azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Azure Credential Error:", err)
	}

	fmt.Println("Azure Blob Connected ✔")
}

// -------------------------------------------------------------
// Upload File to Azure Blob Storage
// -------------------------------------------------------------
func UploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	defer file.Close()

	// Read file
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Create blob name
	ext := filepath.Ext(header.Filename)
	base := strings.TrimSuffix(header.Filename, ext)
	base = strings.ReplaceAll(base, " ", "_")
	blobName := fmt.Sprintf("%s_%d%s", base, time.Now().Unix(), ext)

	// Build blob URL
	blobURL := fmt.Sprintf(
		"https://%s.blob.core.windows.net/%s/%s",
		accountName,
		containerName,
		blobName,
	)

	// Create block blob client using SHARED KEY (correct method for your SDK)
	blobClient, err := blockblob.NewClientWithSharedKeyCredential(blobURL, credential, nil)
	if err != nil {
		return "", fmt.Errorf("blockblob client error: %w", err)
	}

	// Upload file (UploadBuffer exists in your SDK)
	_, err = blobClient.UploadBuffer(
		context.Background(),
		data,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("upload error: %w", err)
	}

	return blobName, nil
}

// -------------------------------------------------------------
// Generate SAS URL
// -------------------------------------------------------------
func GenerateBlobSAS(blobName string) string {
	sig, err := sas.BlobSignatureValues{
		ContainerName: containerName,
		BlobName:      blobName,
		Permissions:   "r",
		ExpiryTime:    time.Now().Add(24 * time.Hour),
	}.SignWithSharedKey(credential)

	if err != nil {
		return ""
	}

	return fmt.Sprintf(
		"https://%s.blob.core.windows.net/%s/%s?%s",
		accountName,
		containerName,
		blobName,
		sig.Encode(),
	)
}
