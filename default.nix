{ pkgs ? import <nixpkgs> { }, buildGoModule ? pkgs.buildGo116Module }:
let
  rev = pkgs.stdenv.mkDerivation {
    name = "rev";
    buildInputs = [ pkgs.git ];
    src = ./.;
    buildPhase = "true";
    installPhase = ''
      echo "$(git rev-parse --short HEAD 2>/dev/null || find * -type f -name '*.go' -print0 | sort -z | xargs -0 sha1sum | sha1sum | sed -r 's/[^\da-f]+//g')" > $out
    '';
  };
  ver = "${pkgs.lib.removeSuffix "\n" (builtins.readFile "${rev}")}";
in buildGoModule {
  pname = "oracle-suite";
  version = "dev-${ver}";
  #  version = pkgs.lib.fileContents ./version;
  src = ./.;
  vendorSha256 = "07d0gkbkz1qnq8bda2vm48s32gg64bv9zicli71qd67nicr6yq51";
  subPackages = [ "cmd/..." ];
  postConfigure = "export CGO_ENABLED=0";
  postInstall = "cp ./config.json $out";
}
