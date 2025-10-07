defmodule Processor.Router do
  use Plug.Router
  require Logger

  plug :match
  plug :dispatch

  get "/health" do
    send_resp(conn, 200, ~s({"status":"ok"}))
  end

  match _ do
    send_resp(conn, 404, "not found")
  end
end

