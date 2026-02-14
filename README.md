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
- Copy `.env.example` to `.env` and fill in your API keys.

#### 1. Ingest News
Fetch and store financial news articles from Alpha Vantage:
- Command: ```go run ./cmd/ingestor {ticker symbol} {lookback # hours}```
- Example: ```go run ./cmd/ingestor NVDA 24```

#### 2. Analyze News
Summarize ingested articles and generate sentiment analysis using Ollama:
- Requires [Ollama](https://ollama.com/) running locally with a model pulled (e.g. `ollama pull llama3.1:8b`)
- Command: ```go run ./cmd/analyzer {lookback # hours}```
- Example: ```go run ./cmd/analyzer 24```

#### Environment Variables
| Variable | Description | Default |
|---|---|---|
| `ALPHAVANTAGE_API_KEY` | Alpha Vantage API key | (required) |
| `OLLAMA_URL` | Ollama server URL | `http://localhost:11434` |
| `OLLAMA_MODEL` | Ollama model to use | `llama3.1:8b` |

### TODO:
- ✅ Fetch
- ✅ Ingest
- ✅ AI Summarization & Sentiment Analysis
- 🚧 Newsletter generation
- 🚧 Email pipeline
