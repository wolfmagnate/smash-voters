# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Running the Application
```bash
go run main.go
```
- Starts the Echo web server on port 8080
- Requires an external research API service running on `http://localhost:8000/research`

### Running with AI Features
```bash
GOOGLE_API_KEY=your_api_key go run main.go
```
- Enables AI-powered logic and debate graph creation
- Required for `/graph` API endpoint functionality

### Building and Testing
```bash
go build
go test ./...
```

### API Testing
```bash
# Health check
curl http://localhost:8080/healthz

# Research API endpoint
curl -X POST "http://localhost:8080/api/v1/research/{theme}/{is_positive}" \
  -H "Content-Type: application/json" \
  -d '{"query": "your research query"}'

# Graph API endpoint
curl -X GET "http://localhost:8080/api/v1/graph?status=completed&result_path=/path/to/result.txt"
curl -X GET "http://localhost:8080/api/v1/graph?status=failed&result_path=/path/to/result.txt"
```

## Architecture Overview

This is a Go web service built with Echo framework that serves as a proxy/wrapper for external research APIs. The application follows a clean architecture pattern with clear separation of concerns.

### Core Components

**HTTP Layer (`handlers/`)**
- `research.go` - Handles POST `/api/v1/research/{theme}/{is_positive}` endpoint
- `graph.go` - Handles GET `/api/v1/graph` endpoint with query parameters
- Validates input, extracts parameters, and delegates to service layer

**Service Layer (`services/`)**
- `research.go` - Business logic for research processing and external API integration
- `graph.go` - AI-powered logic and debate graph creation with Google Generative AI
- Manages HTTP clients and file I/O operations

**Domain Layer (`domain/`)**
- `models/` - Request/response models for API contracts
- `debate_graph.go` - Complex graph structures for debate analysis with nodes, edges, and rebuttals
- `logic_graph.go` - Simpler causal reasoning graph structures
- `rebuttal.go` - Debate rebuttal structures
- `json.go` - JSON utilities

**AI Integration (`infra/`)**
- `ai_client.go` - AI service client integration

**Graph Creation Modules**
- `debate_graph_creator/` - Creates debate analysis graphs from documents
- `logic_graph_creator/` - Creates logical reasoning graphs with cause-effect analysis

### Key Architectural Patterns

1. **Proxy Pattern**: Main service acts as a proxy to external research APIs
2. **Repository Pattern**: Clear separation between domain models and data access
3. **Factory Pattern**: Graph constructors create complex nested structures
4. **Chain of Responsibility**: Multi-step graph analysis pipeline

### Domain Models

The application works with two main graph types:

**Debate Graphs**: Complex argumentation structures with:
- Nodes representing arguments with importance/uniqueness scoring
- Edges representing causal relationships with certainty measures
- Built-in rebuttal mechanisms for counter-arguments

**Logic Graphs**: Simpler cause-effect reasoning chains for logical analysis

### Integration Points

- External research API at `http://localhost:8000/research`
- Google Generative AI integration (`google.golang.org/genai`)
- Echo web framework with CORS, logging, and recovery middleware

### API Endpoints

#### Research API
- **Endpoint**: `POST /api/v1/research/{theme}/{is_positive}`
- **Purpose**: Proxy requests to external research API
- **Parameters**: Path params for theme and stance, JSON body with query

#### Graph API  
- **Endpoint**: `GET /api/v1/graph?status={completed|failed}&result_path={path}`
- **Purpose**: Create AI-powered debate graphs from research results
- **Parameters**: 
  - `status`: "completed" (processes document) or "failed" (returns early)
  - `result_path`: Google Drive path to research result text file
- **Process**: 
  1. Reads research document (currently uses mock positive opinion)
  2. Creates logic graph using AI analysis of argument structure
  3. Enhances with debate graph annotations (importance, uniqueness, rebuttals)
  4. Saves JSON file to `output/` directory with timestamp
- **Requirements**: `GOOGLE_API_KEY` environment variable for AI processing

### File Organization

- AI prompt templates stored as `.md` files alongside corresponding `.go` implementations
- Graph output files saved to `output/` directory with naming pattern: `{path}_debate_graph_{timestamp}.json`
- Clear module separation with domain-driven design
- Component initialization using proper `Create*()` factory functions for AI services