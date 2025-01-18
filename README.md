# YouTube Transcriber API

This is the transcriber used in [Notelytics](https://notelytics.com). It enables transcription of YouTube videos through a simple API interface.

## 🚀 Quick Start

### Pull the Docker Image and Run Locally

Use the pre-built Docker image to get started quickly:
```bash
docker run -d -p 8080:5002 adi555j/youtube-transcriber-go:latest
```

### Build and Run Locally with Docker

If you'd like to build the image yourself:

1. Build the Docker image:
   ```bash
   docker build -t youtube-transcriber-go .
   ```

2. Run the container:
   ```bash
   docker run -d -p 8080:5002 youtube-transcriber-go
   ```

## 📋 API Usage

### Transcription Request Example

To transcribe a YouTube video, send a GET request to the `/transcript` endpoint with the `videoId` parameter:

```bash
curl -X GET "http://localhost:8080/transcript?videoId=dQw4w9WgXcQ"
```

### Parameters
- `videoId`: The ID of the YouTube video to transcribe.

## 🛠 Development

### Prerequisites
- [Docker](https://www.docker.com/) installed on your system.

### Build and Run the Service
1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_name>
   ```

2. Build the Docker image:
   ```bash
   docker build -t youtube-transcriber-go .
   ```

3. Run the container:
   ```bash
   docker run -d -p 8080:5002 youtube-transcriber-go
   ```

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## 📜 License

This project is licensed under the [MIT License](LICENSE).

## ☕ Support

If you find this project useful, consider supporting me by buying me a coffee :). https://buymeacoffee.com/adi555j
---

Happy transcribing!

