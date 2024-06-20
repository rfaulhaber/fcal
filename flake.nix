{
  description = "fcal";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
        };
        name = "fcal";
        version = "1.0.3";
      in {
        packages.fcal = pkgs.buildGoModule {
          inherit version;
          pname = name;
          src = ./.;

          vendorHash = "sha256-+RxuZeH7WFaXrZkBhMTZius0b8DDxUI+kD01dVSOV/k=";
        };

        packages.default = self.packages.${system}.fcal;

        apps.fcal = flake-utils.lib.mkApp {
          drv = self.packages.fcal;
        };
        formatter = pkgs.alejandra;
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [go];
        };
      }
    );
}
