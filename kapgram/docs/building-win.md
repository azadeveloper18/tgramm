# Build Kgramm for Windows

Win64 only. Visual Studio 2026, Windows SDK 10.0.26100.0, toolset `-vcvars_ver=14.44`. The `v145` toolset does not build this tree.

The `api_id` / `api_hash` pair is a client build value. `CheckApiIdAndHash` in the Kgramm server returns nil and does not compare it. Do not pass the published Telegram test pair (`17349` / `344583e45741c457fe1862106095a5eb`). CMake rejects that pair. The shipped default is `api_id` `10000` and `api_hash` `cdd52c2108d4da8b66dffc45aa689a61` (MD5 of the bytes `kgramm`).

## Prepare the prompt

Install Python 3.10 on `PATH` and Git.

```
%comspec% /k "C:\Program Files\Microsoft Visual Studio\18\Community\VC\Auxiliary\Build\vcvars64.bat" -vcvars_ver=14.44
```

Run the rest of this page in that prompt.

## Source and libraries

Pick an empty folder, for example `D:\TBuild`.

```
cd /d D:\TBuild
git clone https://github.com/azadeveloper18/tgramm.git
cd tgramm
git submodule update --init --recursive
kapgram\Telegram\build\prepare\win.bat
```

`prepare\win.bat` writes `Libraries` and `ThirdParty` next to `kapgram` (into the clone). Git submodules of the desktop tree are not stored in this repository. `git submodule update --init --recursive` fetches them from the URLs in `.gitmodules` at the commits recorded for upstream `dev` `5ff615713d2178b8458e201778d23d7abe78db9b`.

## Configure and compile

```
cd /d D:\TBuild\tgramm\kapgram\Telegram
configure.bat x64
```

`x86` and `arm` are rejected. Open `D:\TBuild\tgramm\kapgram\out\Telegram.slnx` in Visual Studio 2026, select the Telegram project, and build Debug and Release. The executable is `kapgram\out\Debug\Kgramm.exe` and `kapgram\out\Release\Kgramm.exe`.

Autoupdate is compiled out. The window title, installer, and icon name this program Kgramm. It connects only to `89.39.95.192` TCP `10443`.
