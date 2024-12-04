# Problem Description

Imagine a scenario where you need to build a system to process large files uploaded by users. The system will read, process, and store metadata from each file. The file processin handle multiple file uploads simultaneg should be scalable toously, and each file's metadata should be stored in a database

### Requirements

1. **Functional Requirements**:

- Accept large file uploads from users.
- Extract and process metadata from each file (e.g., name, size, type, creation date).
- Store the metadata in a database.
- Support concurrent file processing to handle multiple uploads simultaneously.

2. **Non-Functional Requirements**:

- Scalability : Must be able to handle multiple files uploaded simultaneously.
- Reliability: Ensure metadata is accurately extracted and stored.
- Speed: Efficient metadata extraction and storage.

3. **Constraints**:

- File sizes may be large, requiring efficient handling.
- Concurrent processing without data loss or corruption.

### Core Components

- FileUploader: Handles file uploads and manages file storage temporarily for processing.
- FileProcessor: Extracts metadata from files, such as name, size, type, and timestamps.
- MetadataStore: Persists file metadata in a database.
- Queue: Used for concurrent processing of files, enabling scalable processing.

### Data Storage & Models

- Store metadata in a relational or Nosql database
- Metadata fields: FileID, Name, Size, Type, UploadTimestamp.

### Main Operations

- Upload File : Receive and temporarily store the file for processing. Push the file path to a processing queue.

- Process File : Read the file, extract metadata, and generate a FileMetadata object.

- Concurrency : Use a worker pool or goroutines to handle files in the processing queue, enabling multiple files to be processed simultaneously.

