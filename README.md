# Go-Grasscutter

this project is imitate [Grasscutter](https://github.com/Grasscutters/Grasscutter).
It only support for version CN 4.0.

## Current features

- [x] Login
- [ ] Combat
- [ ]  Friends list
- [ ]  Teleportation
- [ ]  Gacha system
- [ ]  Co-op partially works
- [ ]  Spawning monsters via console
- [ ]  Inventory features (receiving items/characters, upgrading items/characters, etc)

### Quick Start (automatic)

- Get [Go 1.22](https://go.dev/dl/)
- Get [MongoDB Community Server](https://www.mongodb.com/try/download/community)
- Get game version CN4.0.x (If you don't have a 4.0.x client, you can find it here and open any of the links to download
  it):
  [4.0.x Client-github](https://github.com/JRSKelvin/GenshinRepository/blob/main/Version%204.0.0.md)
  [4.0.x Client-cloud drive](https://www.123pan.com/s/HoqUVv-U7SBA.html)
- Download the [latest Cultivation version](https://github.com/Grasscutters/Cultivation/releases/latest). Use the `.msi`
  installer.
- After opening Cultivation (as admin), press the download button in the upper right corner.
- Click `Download All-in-One`
- Click the gear in the upper right corner
- Set the game Install path to where your game is located.
- Leave all other settings on default
- Open your build program.
- Click the launch button.
- Log in with whatever username you want. Password can be anything.

### Building

- [Go 1.22](https://go.dev/dl/) or higher
- [Git](https://git-scm.com/downloads)
- [NodeJS](https://nodejs.org/en/download) (Optional, for building the handbook)

##### Clone

```shell
git clone --recurse-submodules https://github.com/issueMel/Go-Grasscutter.git
cd Go-Grasscutter
```

##### Compile

```shell
go build Go-Grasscutter/src
```

##### Compiling the Handbook (Manually)

With Gradle:

```shell
./gradlew generateHandbook
```

With NPM:

```shell
cd src/handbook
npm install
npm run build
```

You can find the output program in the root of the project folder.