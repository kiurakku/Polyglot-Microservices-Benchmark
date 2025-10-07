FROM elixir:1.16-alpine
RUN apk add --no-cache build-base
WORKDIR /app
ENV MIX_ENV=prod
COPY mix.exs mix.lock ./
RUN mix local.hex --force && mix local.rebar --force
RUN mix deps.get
COPY . .
RUN mix compile
EXPOSE 4000
CMD ["elixir", "--erl", "+K true +A 16", "-S", "mix", "run", "--no-halt"]

