#!/bin/bash

echo "Starting tests..."

# Wait for server to start
echo "Waiting for server to start..."
sleep 5  # Adjust this if your server takes longer to start

# Test GetByHeadersHandler
echo "Testing GetByHeadersHandler..."
headers="基本分析分數,技術分析分數,保留盈餘增長标准分数,基因分析標準分數,name"
encodedHeaders=$(echo -n $headers | jq -sRr @uri)
response=$(curl -s "http://localhost:8996/api/v1/StkHK/item?headers=${encodedHeaders}")
echo "Response: ${response}"
echo

# Test GetSingleRecordHandler
echo "Testing GetSingleRecordHandler..."
stockName="1112HK-H&H國際控股"
encodedStockName=$(echo -n $stockName | jq -sRr @uri)
response=$(curl -s "http://localhost:8996/api/v1/StkHK/item/${encodedStockName}")
echo "Response: ${response}"
echo

# Test GetHeadersHandler
echo "Testing GetHeadersHandler..."
response=$(curl -s "http://localhost:8996/api/v1/headers/StkHK")
echo "Response: ${response}"
echo

echo "Finished tests."
