# 运行 speedtest-cli 并将结果保存到变量
speedtest_result=$(speedtest-cli --simple)

# 解析 ping、download 和 upload 数据并构建 JSON 格式
ping=$(echo "$speedtest_result" | awk '/^Ping:/{print $2}')
download=$(echo "$speedtest_result" | awk '/^Download:/{print $2}')
upload=$(echo "$speedtest_result" | awk '/^Upload:/{print $2}')

# 构建 JSON 格式的数据，不包括唯一标识符
data="{\"ping\": \"$ping\", \"download\": \"$download\", \"upload\": \"$upload\"}"

# 使用 cURL 将数据发送到 Elasticsearch，不指定唯一标识符
curl -X POST "http://localhost:9200/your_index/_doc" -H "Content-Type: application/json" -d "$data"
