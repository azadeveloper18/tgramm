/*
This file is part of Telegram Desktop,
the official desktop application for the Telegram messaging service.

For license and copyright information please follow this link:
https://github.com/telegramdesktop/tdesktop/blob/master/LEGAL
*/
#pragma once

#include "base/const_string.h"

#define TDESKTOP_REQUESTED_ALPHA_VERSION (0ULL)

#ifdef TDESKTOP_ALLOW_CLOSED_ALPHA
#define TDESKTOP_ALPHA_VERSION TDESKTOP_REQUESTED_ALPHA_VERSION
#else // TDESKTOP_ALLOW_CLOSED_ALPHA
#define TDESKTOP_ALPHA_VERSION (0ULL)
#endif // TDESKTOP_ALLOW_CLOSED_ALPHA

// Modified by Kgramm on 2026-10-31. See KGRAMM_CHANGES.txt.
// used in Updater.cpp and Setup.iss for Windows
constexpr auto AppId = "{C7FA02C1-4657-56F2-935E-EA4F54CC0D51}"_cs;
constexpr auto AppNameOld = "Kgramm"_cs;
constexpr auto AppName = "Kgramm"_cs;
constexpr auto AppFile = "Kgramm"_cs;
constexpr auto AppVersion = 7003004;
constexpr auto AppVersionStr = "7.3.4";
constexpr auto AppBetaVersion = true;
constexpr auto AppAlphaVersion = TDESKTOP_ALPHA_VERSION;
