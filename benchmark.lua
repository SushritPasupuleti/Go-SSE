wrk.method = "GET"
wrk.body = "{\"from_id\": 1, \"to_id\": 2, \"message\": \"2 followed 1\"}"
wrk.headers["Accept"] = "text/event-stream"

function response()
    os.execute("sleep " .. tonumber(10)) -- replace 10 with the number of seconds you want the request to run
    wrk.thread:stop()
end
