# WhatsApp Bot for Google Sheets Integration

A WhatsApp Business API bot that automatically saves incoming messages to Google Sheets.

## Features

- ✅ WhatsApp webhook verification
- ✅ Parse structured messages and save to Google Sheets
- ✅ Auto-format Google Sheets headers
- ✅ Send confirmation messages back to users
- ✅ Environment-based configuration
- ✅ Clean modular code structure

## Prerequisites

- Go 1.19 or later
- WhatsApp Business API access
- Google Sheets API credentials
- Google Cloud Project with Sheets API enabled

## Setup

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd waba-bot
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup Google Sheets API**
   - Create a Google Cloud Project
   - Enable Google Sheets API
   - Create a service account and download the JSON credentials
   - Save the credentials as `credentials.json` in the project root
   - Share your Google Sheet with the service account email

4. **Setup environment variables**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` with your actual values:
   ```
   VERIFY_TOKEN=your_verify_token
   ACCESS_TOKEN=your_whatsapp_access_token
   PHONE_NUMBER_ID=your_phone_number_id
   SPREADSHEET_ID=your_google_spreadsheet_id
   PORT=3000
   ```

5. **Run the application**
   ```bash
   go run .
   ```

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `VERIFY_TOKEN` | WhatsApp webhook verification token | ✅ |
| `ACCESS_TOKEN` | WhatsApp Business API access token | ✅ |
| `PHONE_NUMBER_ID` | WhatsApp phone number ID | ✅ |
| `SPREADSHEET_ID` | Google Sheets spreadsheet ID | ✅ |
| `PORT` | Server port (default: 3000) | ❌ |

## Message Format

The bot expects messages in the following format:

```
Tanggal: 2025-01-15
Nopol: B1234XYZ
Armada: Truck A
Sopir: John Doe
Pembeli: Company ABC
BB TB: 1000
Bobot: 500
Jenis: Coal
```

## Project Structure

```
├── .env.example          # Environment variables template
├── .gitignore           # Git ignore file
├── README.md            # This file
├── config.go            # Configuration management
├── credentials.json     # Google Sheets API credentials (not in git)
├── go.mod              # Go module file
├── go.sum              # Go dependencies
├── main.go             # Application entry point
├── sheets.go           # Google Sheets service
├── utils.go            # Utility functions
└── whatsapp.go         # WhatsApp webhook handlers
```

## API Endpoints

- `GET /webhook` - WhatsApp webhook verification
- `POST /webhook` - Receive WhatsApp messages

## Development

1. Install Go dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run .
   ```

3. For development with auto-reload, you can use `air`:
   ```bash
   go install github.com/cosmtrek/air@latest
   air
   ```

## Deployment

The application can be deployed to any platform that supports Go applications:

- **Heroku**: Use the included `Procfile`
- **Railway**: Direct deployment from Git
- **Google Cloud Run**: Build and deploy as container
- **Digital Ocean App Platform**: Direct Go deployment

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
# wabb-sheet
