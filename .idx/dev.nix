{ pkgs, ... }: {
  channel = "stable-25.05";
  packages = [
    pkgs.go
    pkgs.nodejs_20
    pkgs.nodePackages.nodemon
  ];
  env = { };
  idx = {
    extensions = [
      "golang.go"
      "google.gemini-cli-vscode-ide-companion"
    ];
    workspace = {
      onCreate = {
        default.openFiles = [ "server.go" ];
      };
    };
    previews = {
      enable = true;
      previews = {
        web = {
          command = [
            "nodemon"
            "--signal"
            "SIGHUP"
            "-w"
            "."
            "-e"
            "go,html"
            "-x"
            "SERVER_ADDRESS=localhost:$PORT go run cmd/mega/main.go"
          ];
          manager = "web";
        };
      };
    };
  };
}
