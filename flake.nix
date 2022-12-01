{
  description = "ZOUT";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    devshell = {
      url = "github:numtide/devshell";
      inputs = {
        flake-utils.follows = "flake-utils";
        nixpkgs.follows = "nixpkgs";
      };
    };
  };

  outputs = { self, nixpkgs, devshell, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; overlays = [ devshell.overlay ]; };
      in
      {
        devShells = rec {
          default = go;
          go = pkgs.devshell.mkShell {
            name = "AOC-go";
            packages = [
              pkgs.nixpkgs-fmt
              pkgs.inotify-tools
              pkgs.go_1_19
              pkgs.go-tools
              pkgs.gotools
            ];
            env = [];
            commands = [];
          };
          elixir = pkgs.devshell.mkShell {
            name = "AOC-elixir";
            packages = [
              pkgs.nixpkgs-fmt
              pkgs.inotify-tools
              pkgs.erlang
              pkgs.elixir
            ];
            env = [];
            commands = [];
          };
          ruby = pkgs.devshell.mkShell {
            name = "AOC-ruby";
            packages = [
              pkgs.nixpkgs-fmt
              pkgs.inotify-tools
              pkgs.ruby
            ];
            env = [];
            commands = [];
          };
          rust = pkgs.devshell.mkShell {
            name = "AOC-rust";
            packages = [
              pkgs.nixpkgs-fmt
              pkgs.inotify-tools
              pkgs.rustup
            ];
            env = [];
            commands = [];
          };
        };
      }
    );
}
