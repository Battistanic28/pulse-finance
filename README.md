# Pulse Finance 🤑

### Description:
An AI data ingest and analysis pipeline for public financial data. Used to generate summaries for an automated newsletter

### Tech Stack:
- Go
- SQLite
- Ollama API
- Alpha Advantage API

### Getting Started:
- First, make sure [Go](https://go.dev/doc/install) is installed on your machine.
- Command: ```go run ./cmd/ingestor {ticker symbol} {lookback # hours}```  
- Example: ```go run ./cmd/ingestor NVDA 24```

### TODO:
- ✅ Fetch
- ✅ Ingest
- 🚧 Dedupe and AI Summarization
- 🚧 AI sentiment analyisis
- 🚧 Newswletter generation
- 🚧 Email pipeline
