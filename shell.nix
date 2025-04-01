{pkgs ? import <nixpkgs> {}}: let
  win2xcur = pkgs.python3Packages.buildPythonPackage rec {
    pname = "win2xcur";
    version = "0.1.2";
    src = pkgs.fetchPypi {
      inherit pname version;
      sha256 = "B8srOXQBUxK6dZ6GhDA5fYvxUBxHVcrSO/z+UWyF+qI=";
    };
    propagatedBuildInputs = with pkgs.python3Packages; [pillow wand];
    doCheck = false;
  };
in
  pkgs.mkShell {
    buildInputs = [
      pkgs.go_1_23
      (pkgs.python3.withPackages (ps: [win2xcur ps.numpy]))
      pkgs.imagemagick
    ];

    shellHook = ''
      echo "Run: go build -o converter converter.go && ./converter"
    '';
  }
