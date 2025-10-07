defmodule Processor.Application do
  use Application
  require Logger

  @impl true
  def start(_type, _args) do
    children = [
      {GNAT, name: :gnat, host: System.get_env("NATS_HOST", "nats"), port: 4222},
      {Task, fn -> subscribe_results() end},
      {Plug.Cowboy, scheme: :http, plug: Processor.Router, options: [port: 4000]}
    ]

    opts = [strategy: :one_for_one, name: Processor.Supervisor]
    Supervisor.start_link(children, opts)
  end

  defp subscribe_results do
    {:ok, conn} = GNAT.whereis(:gnat)
    :ok = GNAT.Sub.start_link(conn, self(), subject: "results")
    loop()
  end

  defp loop do
    receive do
      {:msg, %{body: body}} ->
        Logger.info("Result msg: #{body}")
        loop()
    end
  end
end

