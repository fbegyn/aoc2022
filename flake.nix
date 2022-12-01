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
          default = aoc;
          aoc = pkgs.devshell.mkShell {
            name = "AOC";
            packages = [
              pkgs.nixpkgs-fmt
              pkgs.inotify-tools
              pkgs.go_1_19
              pkgs.gotools
              pkgs.go-tools
              pkgs.erlang
              pkgs.elixir
              pkgs.rustup
            ];
            env = [];
            commands = [];
          };
        };
      }
    );
}
