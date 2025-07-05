#!/bin/bash


echo "🚀 Setting up WhatsApp Bot..."

if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

if [ ! -f .env ]; then
    echo "📝 Creating .env file from template..."
    cp .env.example .env
    echo "⚠️  Please edit .env file with your actual values!"
fi

if [ ! -f credentials.json ]; then
    echo "⚠️  credentials.json not found!"
    echo "   Please download your Google Sheets API credentials and save as 'credentials.json'"
fi

echo "📦 Installing Go dependencies..."
go mod tidy

echo "🔨 Building application..."
go build -o waba-bot

echo "✅ Setup completed!"
echo ""
echo "Next steps:"
echo "1. Edit .env file with your WhatsApp and Google Sheets credentials"
echo "2. Place your Google Sheets API credentials in credentials.json"
echo "3. Run the bot with: ./waba-bot"
echo ""
echo "For development: go run ."
