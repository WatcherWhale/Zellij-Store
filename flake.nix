{
  description = "A way to share environment variables between Zellij Sessions";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      version = "0.0.1";
      rev = "66460e779cd934ff64bfe5c51c5c886c590cdbd4";
      hash = "sha256-WFRMtNXEtktxmYoGj8WrC07xeFoXvEw1mL0sb9M7SPg=";
      vendorHash = "sha256-1F1cdUEp2acf9gxSKA2JLceFys9k5H14QHo4D/TGHwk=";
    in
  {
    packages.x86_64-linux.default = self.packages.x86_64-linux.zellij-store;
    packages.x86_64-linux.zellij-store = nixpkgs.legacyPackages.x86_64-linux.buildGoModule {
      pname = "zellij-store";
      version = "${version}";

      # set to "nixpkgs.lib.fakeHash;", when updating package
      vendorHash = "${vendorHash}";
      proxyVendor = true;
      tags = [ "netgo" ];

      src = nixpkgs.legacyPackages.x86_64-linux.fetchgit {
        name = "zellij-store";
        url = "https://github.com/WatcherWhale/Zellij-Store.git";
        rev = "${rev}";
        hash = "${hash}";
      };

      ldflags = [
        "-w"
        "-s"
      ];

      meta = {
        descirption = "Zellij-Store: A way to share environment variables between Zellij Sessions";
        mainProgram = "zellij-store";
      };

      subPackages = [ "cmd/cli" ];

      postInstall = ''
        mv $out/bin/cli $out/bin/zellij-store
      '';
    };
  };
}
