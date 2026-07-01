# ASCII Art Web: Export File - Zone01

An interactive, production-ready web application built in Go that transforms user-inputted text into stylized graphic ASCII art using distinct typography banner layouts, featuring advanced multi-format document exporting capabilities.

## 🚀 Features
* **Live Generation:** Convert standard English characters and multi-line strings into block-style ASCII visual text.
* **Banner Styles:** Supports three official core typography formats: `Standard`, `Shadow`, and `Thinkertoy`.
* **Dynamic File Exporting:** Seamlessly download generated artwork via a side-by-side structured action system supporting three universal formats:
  * **Plain Text (`.txt`):** Raw unformatted ASCII art character sequences.
  * **Web Page (`.html`):** A portable HTML file with built-in styling, safe dark mode styling, and monospace preservation rules.
  * **API-Ready Data (`.json`):** Formatted structured JSON payloads detailing input string parameters, used banner styles, and the output block.
* **Adaptive File Naming:** Automatically names download bundles based on user choice metrics matching `ascii-art-[banner].[format]` parameters.
* **Safe Input Filtering:** Robust structural validation protecting backend processing channels from irregular inputs or system payloads.
* **Responsive Cyber UI Layout:** Clean, fully integrated CSS theme displaying structural grid assets, error-state handles, and synchronized layouts.

---

## ⚙️ Project Architecture & Design Pattern

The repository relies on a secure, highly modular architectural separation to enforce strict web routing and software design standards:

```text
ascii-art-web-export-file/
├── banners/               # Target layout text fonts (.txt assets)
├── internal/              # Proprietary operational package directories
│   ├── banner/            # Safe text file system asset ingestion routines
│   ├── handlers/          # HTTP request control pipelines, state evaluation, & multi-format export engines
│   ├── render/            # Multi-layer character graphic rendering algorithms
│   ├── server/            # Endpoint initialization & asset distribution routing
│   └── validation/        # Payload structural health constraints
├── static/                # Public assets, browser style rule lists (CSS), & favicons
├── templates/             # Front-end structural templates (HTML environments)
├── cmd/
│   └── main.go            # Central operational system startup entry point
└── go.mod                 # Go operational workspace manifest dependency configuration
```

---

## 🔄 Logical Execution Flow (Order of Operations)

When a client browser triggers an operational request, the software pipeline triggers dependencies down a specific linear workflow:

### 1. Application Initialization (`cmd/main.go`)
* **Role:** The system bootstrapper.
* **Process:** Prints initial telemetry to the host terminal console and signals the proprietary server package to configure communication links.

### 2. Networking and Endpoint Setup (`internal/server/`)
* **Role:** Establishes router mapping models.
* **Process:** Builds standard routing paths via an internal `http.ServeMux`, serves static UI styles from `/static/` directories using file server configurations, and opens continuous web listener loops over port `:8080`.

### 3. Request Orchestration (`internal/handlers/`)
* **Role:** Evaluates method intents and manages file output buffers.
* **Process:** 
  * Shields handlers by blocking unsupported methods (e.g., serving custom `405 Method Not Allowed` layouts for unsafe queries).
  * Evaluates post-generation form action triggers (`generate` vs `export`).
  * If an export action is received, intercepts selection menu parameters, formats outputs down text/HTML/JSON tracks, embeds precise file transfer network instructions, and terminates the response without breaking current client view layouts.

### 4. Payload Interception (`internal/validation/`)
* **Role:** Architectural security firewall.
* **Process:** Scans input boundaries character-by-character to protect text parsing threads from exceptions, restricting entries strictly to printable indexes (ASCII characters 32 to 126) while validating style targets.

### 5. Storage Access (`internal/banner/`)
* **Role:** System file system disk reader.
* **Process:** Locates asset coordinates on disk, handles string loading sequences, and strips out hidden platform carriage returns (`\r`) to guarantee structural integrity across cross-platform instances.

### 6. Typographic Render Engine (`internal/render/`)
* **Role:** Data transformation layer.
* **Process:** Traces character sequences back to line index algorithms (`(char - 32) * 9 + 1`), processes text inputs into multi-layer grid outputs, manages literal code symbols (`\n`), and applies final safety trims.

### 7. Interface Execution (`templates/`)
* **Role:** Client-side workspace compilation.
* **Process:** Dynamically stitches data fields and errors directly into HTML tags, rendering responsive workspaces and side-by-side action controls natively.

---

## 💻 How to Run & Use

### Prerequisites
Make sure you have **Go** installed on your system (version 1.20 or higher recommended).

### 1. Clone the repository
```bash
git clone <your-repository-url>
cd ascii-art-web-export-file
```

### 2. Start the Local Server
Execute the application system from the root directory using the standard entry command:
```bash
go run ./cmd
```

You should see an confirmation terminal diagnostic logging out:
```text
Server starting at http://localhost:8080
```

### 3. Open in Browser
Open your preferred web browser and navigate to:
```text
http://localhost:8080
```

### 4. Generate & Download Artwork
1. Supply target phrases inside the input workspace textarea.
2. Select an active layout font profile (`Standard`, `Shadow`, or `Thinkertoy`).
3. Press **Generate ASCII Art** to render outputs instantly on your dashboard view.
4. Use the dynamic **Format Selector** dropdown to select your target file container (`.txt`, `.html`, or `.json`).
5. Click **Export** to initiate a secure browser file download locally.

---

## 🛠️ Error Codes Standard Map

This project adheres tightly to standard HTTP protocol metrics during verification checks:
* **`200 OK`**: Layout strings resolved cleanly without errors.
* **`400 Bad Request`**: Submissions included non-ASCII entities, unsupported arguments, or corrupt fields.
* **`404 Not Found`**: Request directed to an unregistered path.
* **`405 Method Not Allowed`**: Request targeted endpoints with unsupported HTTP methods.
* **`500 Internal Server Error`**: Core system dependencies or font system resources are missing or broken.

---

## Run Test Files 

Execute granular multi-package unit tests concurrently from your root project level:
```bash
go test ./... -v
```

## 👥 Authors
* **elgeorgiou** - Developer / UI Design / Frontend & Render Specialist
* **gpapadaki** - Developer / System Architecture / Validation & Backend Engineer
