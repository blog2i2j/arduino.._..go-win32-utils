/*
 * This file is part of go-win32-utils.
 *
 * Copyright 2018-2023 ARDUINO SA (http://www.arduino.cc/)
 *
 * go-win32-utils is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 */

package win32

import "syscall"

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall_windows.go

// kernel32.dll

//sys GetModuleHandle(moduleName *byte) (handle syscall.Handle, err error) = GetModuleHandleA

// user32.dll

//sys RegisterClass(wndClass *WndClass) (atom uint16, err error) = user32.RegisterClassA
//sys UnregisterClass(className *byte) (err error) = user32.UnregisterClassA
//sys DefWindowProc(hwnd syscall.Handle, msg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) = user32.DefWindowProcW
//sys CreateWindowEx(exstyle uint32, className *byte, windowText *byte, style uint32, x int32, y int32, width int32, height int32, parent syscall.Handle, menu syscall.Handle, hInstance syscall.Handle, lpParam uintptr) (hwnd syscall.Handle, err error) = user32.CreateWindowExA
//sys DestroyWindowEx(hwnd syscall.Handle) (err error) = user32.DestroyWindow
//sys RegisterDeviceNotification(recipient syscall.Handle, filter *DevBroadcastDeviceInterface, flags uint32) (devHandle syscall.Handle, err error) = user32.RegisterDeviceNotificationA
//sys UnregisterDeviceNotification(deviceHandle syscall.Handle) (err error) = user32.UnregisterDeviceNotification
//sys GetMessage(msg *TagMSG, hwnd syscall.Handle, msgFilterMin uint32, msgFilterMax uint32) (res int32) = user32.GetMessageA
//sys PeekMessage(msg *TagMSG, hwnd syscall.Handle, msgFilterMin uint32, msgFilterMax uint32, removeMsg uint32) (res bool) = user32.PeekMessageA
//sys TranslateMessage(msg *TagMSG) (res bool) = user32.TranslateMessage
//sys DispatchMessage(msg *TagMSG) (res int32) = user32.DispatchMessageA
//sys PostMessage(hwnd syscall.Handle, msg uint32, wParam uintptr, lParam uintptr) (res bool) = user32.PostMessageA

// shell32.dll

//sys getKnownFolderPath(rfid *syscall.GUID, dwFlags uint32, hToken syscall.Handle, path **uint16) (regerrno error) = shell32.SHGetKnownFolderPath
//sys getFolderPath(hwndOwner uint32, nFolder int, hToken syscall.Handle, dwFlags uint32, path *uint16) (regerrno error) = shell32.SHGetFolderPathW

// ole32.dll

//sys taskMemFree(pv uintptr) = ole32.CoTaskMemFree

type folderIdentifier struct {
	FOLDERID *syscall.GUID
	CSIDL    int
}

// Windows folderID constants
var folderIDAddNewPrograms = &syscall.GUID{Data1: 0xDE61D971, Data2: 0x5EBC, Data3: 0x4F02, Data4: [8]byte{0xA3, 0xA9, 0x6C, 0x82, 0x89, 0x5E, 0x5C, 0x04}}
var folderIDAdminTools = &syscall.GUID{Data1: 0x724EF170, Data2: 0xA42D, Data3: 0x4FEF, Data4: [8]byte{0x9F, 0x26, 0xB6, 0x0E, 0x84, 0x6F, 0xBA, 0x4F}}
var folderIDAppUpdates = &syscall.GUID{Data1: 0xA305CE99, Data2: 0xF527, Data3: 0x492B, Data4: [8]byte{0x8B, 0x1A, 0x7E, 0x76, 0xFA, 0x98, 0xD6, 0xE4}}
var folderIDCDBurning = &syscall.GUID{Data1: 0x9E52AB10, Data2: 0xF80D, Data3: 0x49DF, Data4: [8]byte{0xAC, 0xB8, 0x43, 0x30, 0xF5, 0x68, 0x78, 0x55}}
var folderIDChangeRemovePrograms = &syscall.GUID{Data1: 0xDF7266AC, Data2: 0x9274, Data3: 0x4867, Data4: [8]byte{0x8D, 0x55, 0x3B, 0xD6, 0x61, 0xDE, 0x87, 0x2D}}
var folderIDCommonAdminTools = &syscall.GUID{Data1: 0xD0384E7D, Data2: 0xBAC3, Data3: 0x4797, Data4: [8]byte{0x8F, 0x14, 0xCB, 0xA2, 0x29, 0xB3, 0x92, 0xB5}}
var folderIDCommonOEMLinks = &syscall.GUID{Data1: 0xC1BAE2D0, Data2: 0x10DF, Data3: 0x4334, Data4: [8]byte{0xBE, 0xDD, 0x7A, 0xA2, 0x0B, 0x22, 0x7A, 0x9D}}
var folderIDCommonPrograms = &syscall.GUID{Data1: 0x0139D44E, Data2: 0x6AFE, Data3: 0x49F2, Data4: [8]byte{0x86, 0x90, 0x3D, 0xAF, 0xCA, 0xE6, 0xFF, 0xB8}}
var folderIDCommonStartMenu = &syscall.GUID{Data1: 0xA4115719, Data2: 0xD62E, Data3: 0x491D, Data4: [8]byte{0xAA, 0x7C, 0xE7, 0x4B, 0x8B, 0xE3, 0xB0, 0x67}}
var folderIDCommonStartup = &syscall.GUID{Data1: 0x82A5EA35, Data2: 0xD9CD, Data3: 0x47C5, Data4: [8]byte{0x96, 0x29, 0xE1, 0x5D, 0x2F, 0x71, 0x4E, 0x6E}}
var folderIDCommonTemplates = &syscall.GUID{Data1: 0xB94237E7, Data2: 0x57AC, Data3: 0x4347, Data4: [8]byte{0x91, 0x51, 0xB0, 0x8C, 0x6C, 0x32, 0xD1, 0xF7}}
var folderIDComputerFolder = &syscall.GUID{Data1: 0x0AC0837C, Data2: 0xBBF8, Data3: 0x452A, Data4: [8]byte{0x85, 0x0D, 0x79, 0xD0, 0x8E, 0x66, 0x7C, 0xA7}}
var folderIDConflictFolder = &syscall.GUID{Data1: 0x4BFEFB45, Data2: 0x347D, Data3: 0x4006, Data4: [8]byte{0xA5, 0xBE, 0xAC, 0x0C, 0xB0, 0x56, 0x71, 0x92}}
var folderIDConnectionsFolder = &syscall.GUID{Data1: 0x6F0CD92B, Data2: 0x2E97, Data3: 0x45D1, Data4: [8]byte{0x88, 0xFF, 0xB0, 0xD1, 0x86, 0xB8, 0xDE, 0xDD}}
var folderIDContacts = &syscall.GUID{Data1: 0x56784854, Data2: 0xC6CB, Data3: 0x462B, Data4: [8]byte{0x81, 0x69, 0x88, 0xE3, 0x50, 0xAC, 0xB8, 0x82}}
var folderIDControlPanelFolder = &syscall.GUID{Data1: 0x82A74AEB, Data2: 0xAEB4, Data3: 0x465C, Data4: [8]byte{0xA0, 0x14, 0xD0, 0x97, 0xEE, 0x34, 0x6D, 0x63}}
var folderIDCookies = &syscall.GUID{Data1: 0x2B0F765D, Data2: 0xC0E9, Data3: 0x4171, Data4: [8]byte{0x90, 0x8E, 0x08, 0xA6, 0x11, 0xB8, 0x4F, 0xF6}}
var folderIDDesktop = &syscall.GUID{Data1: 0xB4BFCC3A, Data2: 0xDB2C, Data3: 0x424C, Data4: [8]byte{0xB0, 0x29, 0x7F, 0xE9, 0x9A, 0x87, 0xC6, 0x41}}
var folderIDDeviceMetadataStore = &syscall.GUID{Data1: 0x5CE4A5E9, Data2: 0xE4EB, Data3: 0x479D, Data4: [8]byte{0xB8, 0x9F, 0x13, 0x0C, 0x02, 0x88, 0x61, 0x55}}
var folderIDDocuments = &syscall.GUID{Data1: 0xFDD39AD0, Data2: 0x238F, Data3: 0x46AF, Data4: [8]byte{0xAD, 0xB4, 0x6C, 0x85, 0x48, 0x03, 0x69, 0xC7}}
var folderIDDocumentsLibrary = &syscall.GUID{Data1: 0x7B0DB17D, Data2: 0x9CD2, Data3: 0x4A93, Data4: [8]byte{0x97, 0x33, 0x46, 0xCC, 0x89, 0x02, 0x2E, 0x7C}}
var folderIDDownloads = &syscall.GUID{Data1: 0x374DE290, Data2: 0x123F, Data3: 0x4565, Data4: [8]byte{0x91, 0x64, 0x39, 0xC4, 0x92, 0x5E, 0x46, 0x7B}}
var folderIDFavorites = &syscall.GUID{Data1: 0x1777F761, Data2: 0x68AD, Data3: 0x4D8A, Data4: [8]byte{0x87, 0xBD, 0x30, 0xB7, 0x59, 0xFA, 0x33, 0xDD}}
var folderIDFonts = &syscall.GUID{Data1: 0xFD228CB7, Data2: 0xAE11, Data3: 0x4AE3, Data4: [8]byte{0x86, 0x4C, 0x16, 0xF3, 0x91, 0x0A, 0xB8, 0xFE}}
var folderIDGames = &syscall.GUID{Data1: 0xCAC52C1A, Data2: 0xB53D, Data3: 0x4EDC, Data4: [8]byte{0x92, 0xD7, 0x6B, 0x2E, 0x8A, 0xC1, 0x94, 0x34}}
var folderIDGameTasks = &syscall.GUID{Data1: 0x054FAE61, Data2: 0x4DD8, Data3: 0x4787, Data4: [8]byte{0x80, 0xB6, 0x09, 0x02, 0x20, 0xC4, 0xB7, 0x00}}
var folderIDHistory = &syscall.GUID{Data1: 0xD9DC8A3B, Data2: 0xB784, Data3: 0x432E, Data4: [8]byte{0xA7, 0x81, 0x5A, 0x11, 0x30, 0xA7, 0x59, 0x63}}
var folderIDHomeGroup = &syscall.GUID{Data1: 0x52528A6B, Data2: 0xB9E3, Data3: 0x4ADD, Data4: [8]byte{0xB6, 0x0D, 0x58, 0x8C, 0x2D, 0xBA, 0x84, 0x2D}}
var folderIDImplicitAppShortcuts = &syscall.GUID{Data1: 0xBCB5256F, Data2: 0x79F6, Data3: 0x4CEE, Data4: [8]byte{0xB7, 0x25, 0xDC, 0x34, 0xE4, 0x02, 0xFD, 0x46}}
var folderIDInternetCache = &syscall.GUID{Data1: 0x352481E8, Data2: 0x33BE, Data3: 0x4251, Data4: [8]byte{0xBA, 0x85, 0x60, 0x07, 0xCA, 0xED, 0xCF, 0x9D}}
var folderIDInternetFolder = &syscall.GUID{Data1: 0x4D9F7874, Data2: 0x4E0C, Data3: 0x4904, Data4: [8]byte{0x96, 0x7B, 0x40, 0xB0, 0xD2, 0x0C, 0x3E, 0x4B}}
var folderIDLibraries = &syscall.GUID{Data1: 0x1B3EA5DC, Data2: 0xB587, Data3: 0x4786, Data4: [8]byte{0xB4, 0xEF, 0xBD, 0x1D, 0xC3, 0x32, 0xAE, 0xAE}}
var folderIDLinks = &syscall.GUID{Data1: 0xBFB9D5E0, Data2: 0xC6A9, Data3: 0x404C, Data4: [8]byte{0xB2, 0xB2, 0xAE, 0x6D, 0xB6, 0xAF, 0x49, 0x68}}
var folderIDLocalAppData = &syscall.GUID{Data1: 0xF1B32785, Data2: 0x6FBA, Data3: 0x4FCF, Data4: [8]byte{0x9D, 0x55, 0x7B, 0x8E, 0x7F, 0x15, 0x70, 0x91}}
var folderIDLocalAppDataLow = &syscall.GUID{Data1: 0xA520A1A4, Data2: 0x1780, Data3: 0x4FF6, Data4: [8]byte{0xBD, 0x18, 0x16, 0x73, 0x43, 0xC5, 0xAF, 0x16}}
var folderIDLocalizedResourcesDir = &syscall.GUID{Data1: 0x2A00375E, Data2: 0x224C, Data3: 0x49DE, Data4: [8]byte{0xB8, 0xD1, 0x44, 0x0D, 0xF7, 0xEF, 0x3D, 0xDC}}
var folderIDMusic = &syscall.GUID{Data1: 0x4BD8D571, Data2: 0x6D19, Data3: 0x48D3, Data4: [8]byte{0xBE, 0x97, 0x42, 0x22, 0x20, 0x08, 0x0E, 0x43}}
var folderIDMusicLibrary = &syscall.GUID{Data1: 0x2112AB0A, Data2: 0xC86A, Data3: 0x4FFE, Data4: [8]byte{0xA3, 0x68, 0x0D, 0xE9, 0x6E, 0x47, 0x01, 0x2E}}
var folderIDNetHood = &syscall.GUID{Data1: 0xC5ABBF53, Data2: 0xE17F, Data3: 0x4121, Data4: [8]byte{0x89, 0x00, 0x86, 0x62, 0x6F, 0xC2, 0xC9, 0x73}}
var folderIDNetworkFolder = &syscall.GUID{Data1: 0xD20BEEC4, Data2: 0x5CA8, Data3: 0x4905, Data4: [8]byte{0xAE, 0x3B, 0xBF, 0x25, 0x1E, 0xA0, 0x9B, 0x53}}
var folderIDOriginalImages = &syscall.GUID{Data1: 0x2C36C0AA, Data2: 0x5812, Data3: 0x4B87, Data4: [8]byte{0xBF, 0xD0, 0x4C, 0xD0, 0xDF, 0xB1, 0x9B, 0x39}}
var folderIDPhotoAlbums = &syscall.GUID{Data1: 0x69D2CF90, Data2: 0xFC33, Data3: 0x4FB7, Data4: [8]byte{0x9A, 0x0C, 0xEB, 0xB0, 0xF0, 0xFC, 0xB4, 0x3C}}
var folderIDPictures = &syscall.GUID{Data1: 0x33E28130, Data2: 0x4E1E, Data3: 0x4676, Data4: [8]byte{0x83, 0x5A, 0x98, 0x39, 0x5C, 0x3B, 0xC3, 0xBB}}
var folderIDPicturesLibrary = &syscall.GUID{Data1: 0xA990AE9F, Data2: 0xA03B, Data3: 0x4E80, Data4: [8]byte{0x94, 0xBC, 0x99, 0x12, 0xD7, 0x50, 0x41, 0x04}}
var folderIDPlaylists = &syscall.GUID{Data1: 0xDE92C1C7, Data2: 0x837F, Data3: 0x4F69, Data4: [8]byte{0xA3, 0xBB, 0x86, 0xE6, 0x31, 0x20, 0x4A, 0x23}}
var folderIDPrintersFolder = &syscall.GUID{Data1: 0x76FC4E2D, Data2: 0xD6AD, Data3: 0x4519, Data4: [8]byte{0xA6, 0x63, 0x37, 0xBD, 0x56, 0x06, 0x81, 0x85}}
var folderIDPrintHood = &syscall.GUID{Data1: 0x9274BD8D, Data2: 0xCFD1, Data3: 0x41C3, Data4: [8]byte{0xB3, 0x5E, 0xB1, 0x3F, 0x55, 0xA7, 0x58, 0xF4}}
var folderIDProfile = &syscall.GUID{Data1: 0x5E6C858F, Data2: 0x0E22, Data3: 0x4760, Data4: [8]byte{0x9A, 0xFE, 0xEA, 0x33, 0x17, 0xB6, 0x71, 0x73}}
var folderIDProgramData = &syscall.GUID{Data1: 0x62AB5D82, Data2: 0xFDC1, Data3: 0x4DC3, Data4: [8]byte{0xA9, 0xDD, 0x07, 0x0D, 0x1D, 0x49, 0x5D, 0x97}}
var folderIDProgramFiles = &syscall.GUID{Data1: 0x905E63B6, Data2: 0xC1BF, Data3: 0x494E, Data4: [8]byte{0xB2, 0x9C, 0x65, 0xB7, 0x32, 0xD3, 0xD2, 0x1A}}
var folderIDProgramFilesCommon = &syscall.GUID{Data1: 0xF7F1ED05, Data2: 0x9F6D, Data3: 0x47A2, Data4: [8]byte{0xAA, 0xAE, 0x29, 0xD3, 0x17, 0xC6, 0xF0, 0x66}}
var folderIDProgramFilesCommonX64 = &syscall.GUID{Data1: 0x6365D5A7, Data2: 0x0F0D, Data3: 0x45E5, Data4: [8]byte{0x87, 0xF6, 0x0D, 0xA5, 0x6B, 0x6A, 0x4F, 0x7D}}
var folderIDProgramFilesCommonX86 = &syscall.GUID{Data1: 0xDE974D24, Data2: 0xD9C6, Data3: 0x4D3E, Data4: [8]byte{0xBF, 0x91, 0xF4, 0x45, 0x51, 0x20, 0xB9, 0x17}}
var folderIDProgramFilesX64 = &syscall.GUID{Data1: 0x6D809377, Data2: 0x6AF0, Data3: 0x444B, Data4: [8]byte{0x89, 0x57, 0xA3, 0x77, 0x3F, 0x02, 0x20, 0x0E}}
var folderIDProgramFilesX86 = &syscall.GUID{Data1: 0x7C5A40EF, Data2: 0xA0FB, Data3: 0x4BFC, Data4: [8]byte{0x87, 0x4A, 0xC0, 0xF2, 0xE0, 0xB9, 0xFA, 0x8E}}
var folderIDPrograms = &syscall.GUID{Data1: 0xA77F5D77, Data2: 0x2E2B, Data3: 0x44C3, Data4: [8]byte{0xA6, 0xA2, 0xAB, 0xA6, 0x01, 0x05, 0x4A, 0x51}}
var folderIDPublic = &syscall.GUID{Data1: 0xDFDF76A2, Data2: 0xC82A, Data3: 0x4D63, Data4: [8]byte{0x90, 0x6A, 0x56, 0x44, 0xAC, 0x45, 0x73, 0x85}}
var folderIDPublicDesktop = &syscall.GUID{Data1: 0xC4AA340D, Data2: 0xF20F, Data3: 0x4863, Data4: [8]byte{0xAF, 0xEF, 0xF8, 0x7E, 0xF2, 0xE6, 0xBA, 0x25}}
var folderIDPublicDocuments = &syscall.GUID{Data1: 0xED4824AF, Data2: 0xDCE4, Data3: 0x45A8, Data4: [8]byte{0x81, 0xE2, 0xFC, 0x79, 0x65, 0x08, 0x36, 0x34}}
var folderIDPublicDownloads = &syscall.GUID{Data1: 0x3D644C9B, Data2: 0x1FB8, Data3: 0x4F30, Data4: [8]byte{0x9B, 0x45, 0xF6, 0x70, 0x23, 0x5F, 0x79, 0xC0}}
var folderIDPublicGameTasks = &syscall.GUID{Data1: 0xDEBF2536, Data2: 0xE1A8, Data3: 0x4C59, Data4: [8]byte{0xB6, 0xA2, 0x41, 0x45, 0x86, 0x47, 0x6A, 0xEA}}
var folderIDPublicLibraries = &syscall.GUID{Data1: 0x48DAF80B, Data2: 0xE6CF, Data3: 0x4F4E, Data4: [8]byte{0xB8, 0x00, 0x0E, 0x69, 0xD8, 0x4E, 0xE3, 0x84}}
var folderIDPublicMusic = &syscall.GUID{Data1: 0x3214FAB5, Data2: 0x9757, Data3: 0x4298, Data4: [8]byte{0xBB, 0x61, 0x92, 0xA9, 0xDE, 0xAA, 0x44, 0xFF}}
var folderIDPublicPictures = &syscall.GUID{Data1: 0xB6EBFB86, Data2: 0x6907, Data3: 0x413C, Data4: [8]byte{0x9A, 0xF7, 0x4F, 0xC2, 0xAB, 0xF0, 0x7C, 0xC5}}
var folderIDPublicRingtones = &syscall.GUID{Data1: 0xE555AB60, Data2: 0x153B, Data3: 0x4D17, Data4: [8]byte{0x9F, 0x04, 0xA5, 0xFE, 0x99, 0xFC, 0x15, 0xEC}}
var folderIDPublicVideos = &syscall.GUID{Data1: 0x2400183A, Data2: 0x6185, Data3: 0x49FB, Data4: [8]byte{0xA2, 0xD8, 0x4A, 0x39, 0x2A, 0x60, 0x2B, 0xA3}}
var folderIDQuickLaunch = &syscall.GUID{Data1: 0x52A4F021, Data2: 0x7B75, Data3: 0x48A9, Data4: [8]byte{0x9F, 0x6B, 0x4B, 0x87, 0xA2, 0x10, 0xBC, 0x8F}}
var folderIDRecent = &syscall.GUID{Data1: 0xAE50C081, Data2: 0xEBD2, Data3: 0x438A, Data4: [8]byte{0x86, 0x55, 0x8A, 0x09, 0x2E, 0x34, 0x98, 0x7A}}
var folderIDRecordedTVLibrary = &syscall.GUID{Data1: 0x1A6FDBA2, Data2: 0xF42D, Data3: 0x4358, Data4: [8]byte{0xA7, 0x98, 0xB7, 0x4D, 0x74, 0x59, 0x26, 0xC5}}
var folderIDRecycleBinFolder = &syscall.GUID{Data1: 0xB7534046, Data2: 0x3ECB, Data3: 0x4C18, Data4: [8]byte{0xBE, 0x4E, 0x64, 0xCD, 0x4C, 0xB7, 0xD6, 0xAC}}
var folderIDResourceDir = &syscall.GUID{Data1: 0x8AD10C31, Data2: 0x2ADB, Data3: 0x4296, Data4: [8]byte{0xA8, 0xF7, 0xE4, 0x70, 0x12, 0x32, 0xC9, 0x72}}
var folderIDRingtones = &syscall.GUID{Data1: 0xC870044B, Data2: 0xF49E, Data3: 0x4126, Data4: [8]byte{0xA9, 0xC3, 0xB5, 0x2A, 0x1F, 0xF4, 0x11, 0xE8}}
var folderIDRoamingAppData = &syscall.GUID{Data1: 0x3EB685DB, Data2: 0x65F9, Data3: 0x4CF6, Data4: [8]byte{0xA0, 0x3A, 0xE3, 0xEF, 0x65, 0x72, 0x9F, 0x3D}}
var folderIDSampleMusic = &syscall.GUID{Data1: 0xB250C668, Data2: 0xF57D, Data3: 0x4EE1, Data4: [8]byte{0xA6, 0x3C, 0x29, 0x0E, 0xE7, 0xD1, 0xAA, 0x1F}}
var folderIDSamplePictures = &syscall.GUID{Data1: 0xC4900540, Data2: 0x2379, Data3: 0x4C75, Data4: [8]byte{0x84, 0x4B, 0x64, 0xE6, 0xFA, 0xF8, 0x71, 0x6B}}
var folderIDSamplePlaylists = &syscall.GUID{Data1: 0x15CA69B3, Data2: 0x30EE, Data3: 0x49C1, Data4: [8]byte{0xAC, 0xE1, 0x6B, 0x5E, 0xC3, 0x72, 0xAF, 0xB5}}
var folderIDSampleVideos = &syscall.GUID{Data1: 0x859EAD94, Data2: 0x2E85, Data3: 0x48AD, Data4: [8]byte{0xA7, 0x1A, 0x09, 0x69, 0xCB, 0x56, 0xA6, 0xCD}}
var folderIDSavedGames = &syscall.GUID{Data1: 0x4C5C32FF, Data2: 0xBB9D, Data3: 0x43B0, Data4: [8]byte{0xB5, 0xB4, 0x2D, 0x72, 0xE5, 0x4E, 0xAA, 0xA4}}
var folderIDSavedSearches = &syscall.GUID{Data1: 0x7D1D3A04, Data2: 0xDEBB, Data3: 0x4115, Data4: [8]byte{0x95, 0xCF, 0x2F, 0x29, 0xDA, 0x29, 0x20, 0xDA}}
var folderIDSearchHome = &syscall.GUID{Data1: 0x190337D1, Data2: 0xB8CA, Data3: 0x4121, Data4: [8]byte{0xA6, 0x39, 0x6D, 0x47, 0x2D, 0x16, 0x97, 0x2A}}
var folderIDSearchCSC = &syscall.GUID{Data1: 0xEE32E446, Data2: 0x31CA, Data3: 0x4ABA, Data4: [8]byte{0x81, 0x4F, 0xA5, 0xEB, 0xD2, 0xFD, 0x6D, 0x5E}}
var folderIDSearchMAPI = &syscall.GUID{Data1: 0x98EC0E18, Data2: 0x2098, Data3: 0x4D44, Data4: [8]byte{0x86, 0x44, 0x66, 0x97, 0x93, 0x15, 0xA2, 0x81}}
var folderIDSendTo = &syscall.GUID{Data1: 0x8983036C, Data2: 0x27C0, Data3: 0x404B, Data4: [8]byte{0x8F, 0x08, 0x10, 0x2D, 0x10, 0xDC, 0xFD, 0x74}}
var folderIDSidebarDefaultParts = &syscall.GUID{Data1: 0x7B396E54, Data2: 0x9EC5, Data3: 0x4300, Data4: [8]byte{0xBE, 0x0A, 0x24, 0x82, 0xEB, 0xAE, 0x1A, 0x26}}
var folderIDSidebarParts = &syscall.GUID{Data1: 0xA75D362E, Data2: 0x50FC, Data3: 0x4FB7, Data4: [8]byte{0xAC, 0x2C, 0xA8, 0xBE, 0xAA, 0x31, 0x44, 0x93}}
var folderIDStartMenu = &syscall.GUID{Data1: 0x625B53C3, Data2: 0xAB48, Data3: 0x4EC1, Data4: [8]byte{0xBA, 0x1F, 0xA1, 0xEF, 0x41, 0x46, 0xFC, 0x19}}
var folderIDStartup = &syscall.GUID{Data1: 0xB97D20BB, Data2: 0xF46A, Data3: 0x4C97, Data4: [8]byte{0xBA, 0x10, 0x5E, 0x36, 0x08, 0x43, 0x08, 0x54}}
var folderIDSyncManagerFolder = &syscall.GUID{Data1: 0x43668BF8, Data2: 0xC14E, Data3: 0x49B2, Data4: [8]byte{0x97, 0xC9, 0x74, 0x77, 0x84, 0xD7, 0x84, 0xB7}}
var folderIDSyncResultsFolder = &syscall.GUID{Data1: 0x289A9A43, Data2: 0xBE44, Data3: 0x4057, Data4: [8]byte{0xA4, 0x1B, 0x58, 0x7A, 0x76, 0xD7, 0xE7, 0xF9}}
var folderIDSyncSetupFolder = &syscall.GUID{Data1: 0x0F214138, Data2: 0xB1D3, Data3: 0x4A90, Data4: [8]byte{0xBB, 0xA9, 0x27, 0xCB, 0xC0, 0xC5, 0x38, 0x9A}}
var folderIDSystem = &syscall.GUID{Data1: 0x1AC14E77, Data2: 0x02E7, Data3: 0x4E5D, Data4: [8]byte{0xB7, 0x44, 0x2E, 0xB1, 0xAE, 0x51, 0x98, 0xB7}}
var folderIDSystemX86 = &syscall.GUID{Data1: 0xD65231B0, Data2: 0xB2F1, Data3: 0x4857, Data4: [8]byte{0xA4, 0xCE, 0xA8, 0xE7, 0xC6, 0xEA, 0x7D, 0x27}}
var folderIDTemplates = &syscall.GUID{Data1: 0xA63293E8, Data2: 0x664E, Data3: 0x48DB, Data4: [8]byte{0xA0, 0x79, 0xDF, 0x75, 0x9E, 0x05, 0x09, 0xF7}}
var folderIDUserPinned = &syscall.GUID{Data1: 0x9E3995AB, Data2: 0x1F9C, Data3: 0x4F13, Data4: [8]byte{0xB8, 0x27, 0x48, 0xB2, 0x4B, 0x6C, 0x71, 0x74}}
var folderIDUserProfiles = &syscall.GUID{Data1: 0x0762D272, Data2: 0xC50A, Data3: 0x4BB0, Data4: [8]byte{0xA3, 0x82, 0x69, 0x7D, 0xCD, 0x72, 0x9B, 0x80}}
var folderIDUserProgramFiles = &syscall.GUID{Data1: 0x5CD7AEE2, Data2: 0x2219, Data3: 0x4A67, Data4: [8]byte{0xB8, 0x5D, 0x6C, 0x9C, 0xE1, 0x56, 0x60, 0xCB}}
var folderIDUserProgramFilesCommon = &syscall.GUID{Data1: 0xBCBD3057, Data2: 0xCA5C, Data3: 0x4622, Data4: [8]byte{0xB4, 0x2D, 0xBC, 0x56, 0xDB, 0x0A, 0xE5, 0x16}}
var folderIDUsersFiles = &syscall.GUID{Data1: 0xF3CE0F7C, Data2: 0x4901, Data3: 0x4ACC, Data4: [8]byte{0x86, 0x48, 0xD5, 0xD4, 0x4B, 0x04, 0xEF, 0x8F}}
var folderIDUsersLibraries = &syscall.GUID{Data1: 0xA302545D, Data2: 0xDEFF, Data3: 0x464B, Data4: [8]byte{0xAB, 0xE8, 0x61, 0xC8, 0x64, 0x8D, 0x93, 0x9B}}
var folderIDVideos = &syscall.GUID{Data1: 0x18989B1D, Data2: 0x99B5, Data3: 0x455B, Data4: [8]byte{0x84, 0x1C, 0xAB, 0x7C, 0x74, 0xE4, 0xDD, 0xFC}}
var folderIDVideosLibrary = &syscall.GUID{Data1: 0x491E922F, Data2: 0x5643, Data3: 0x4AF4, Data4: [8]byte{0xA7, 0xEB, 0x4E, 0x7A, 0x13, 0x8D, 0x81, 0x74}}
var folderIDWindows = &syscall.GUID{Data1: 0xF38BF404, Data2: 0x1D43, Data3: 0x42F2, Data4: [8]byte{0x93, 0x05, 0x67, 0xDE, 0x0B, 0x28, 0xFC, 0x23}}

// Windows CSIDL constants
const csidlAdminTools = 48
const csidlAltStartup = 29
const csidlAppData = 26
const csidlBitBucket = 10
const csidlCdBurnArea = 59
const csidlCommonAdminTools = 47
const csidlCommonAltStartup = 30
const csidlCommonAppData = 35
const csidlCommonDesktopDirectory = 25
const csidlCommonDocuments = 46
const csidlCommonFavorites = 31
const csidlCommonMusic = 53
const csidlCommonPictures = 54
const csidlCommonPrograms = 23
const csidlCommonStartMenu = 22
const csidlCommonStartup = 24
const csidlCommonTemplates = 45
const csidlCommonVideo = 55
const csidlComputersNearMe = 61
const csidlConnections = 49
const csidlControls = 3
const csidlCookies = 33
const csidlDesktop = 0
const csidlDesktopDirectory = 16
const csidlDrives = 17
const csidlFavorites = 6
const csidlFonts = 20
const csidlHistory = 34
const csidlInternet = 1
const csidlInternetCache = 32
const csidlLolcaAppData = 28
const csidlMyDocuments = 5
const csidlMyMusic = 13
const csidlMyPictures = 39
const csidlMyVideo = 14
const csidlNethood = 19
const csidlNetwork = 18
const csidlPersonal = 5
const csidlPhotoAlbums = 69
const csidlPLAYLISTS = 63
const csidlPrinters = 4
const csidlPrinthood = 27
const csidlProfile = 40
const csidlProgramFiles = 38
const csidlProgramFilesX86 = 42
const csidlProgramFilesCommon = 43
const csidlProgramFilesCommonX86 = 44
const csidlPrograms = 2
const csidlRecent = 8
const csidlResources = 56
const csidlResourcesLocalized = 57
const csidlSampleMusic = 64
const csidlSamplePlaylists = 65
const csidlSamplePicrtures = 66
const csidlSampleVideos = 67
const csidlSendTo = 9
const csidlStartMenu = 11
const csidlStartup = 7
const csidlSystem = 37
const csidlSystemX86 = 41
const csidlTemplates = 21
const csidlWindows = 36

// WndClass FIXMEDOCS
type WndClass struct {
	Style        uint32
	WndProc      uintptr
	ClsExtra     int32
	WndExtra     int32
	Instance     syscall.Handle
	Icon         syscall.Handle
	Cursor       syscall.Handle
	BrBackground syscall.Handle
	MenuName     *byte
	ClassName    *byte
}

// Point FIXMEDOCS
type Point struct {
	X int32
	Y int32
}

// TagMSG FIXMEDOCS
type TagMSG struct {
	Hwnd     syscall.Handle
	Message  uint32
	WParam   uintptr
	LParam   uintptr
	Time     int32
	Pt       Point
	LPrivate int32
}

// WMQuit FIXMEDOCS
const WMQuit = 0x0012

const (
	// WsExDlgModalFrame FIXMEDOCS
	WsExDlgModalFrame = 0x00000001
	// WsExTopmost FIXMEDOCS
	WsExTopmost = 0x00000008
	// WsExTransparent FIXMEDOCS
	WsExTransparent = 0x00000020
	// WsExMDIChild FIXMEDOCS
	WsExMDIChild = 0x00000040
	// WsExToolWindow FIXMEDOCS
	WsExToolWindow = 0x00000080
	// WsExAppWindow FIXMEDOCS
	WsExAppWindow = 0x00040000
	// WsExLayered FIXMEDOCS
	WsExLayered = 0x00080000
)

// GUID FIXEMEDOCS
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// DevBroadcastDeviceInterface FIXMEDOCS
type DevBroadcastDeviceInterface struct {
	DwSize       uint32
	DwDeviceType uint32
	DwReserved   uint32
	ClassGUID    GUID
	SzName       uint16
}

// UsbEventGUID is USB devices GUID used to filter notifications
var UsbEventGUID GUID = GUID{
	Data1: 0x10bfdca5,
	Data2: 0x3065,
	Data3: 0xd211,
	Data4: [8]byte{0x90, 0x1f, 0x00, 0xc0, 0x4f, 0xb9, 0x51, 0xed},
}

const (
	// DeviceNotifyWindowHandle FIXMEDOCS
	DeviceNotifyWindowHandle = 0
	// DeviceNotifySserviceHandle FIXMEDOCS
	DeviceNotifySserviceHandle = 1
	// DeviceNotifyAllInterfaceClasses FIXMEDOCS
	DeviceNotifyAllInterfaceClasses = 4
)

// DbtDevtypeDeviceInterface FIXMEDOCS
const DbtDevtypeDeviceInterface = 5

const (
	// PMNoRemove FIXMEDOCS
	PMNoRemove = 0x0000
	// PMRemove FIXMEDOCS
	PMRemove = 0x0001
	// PMNoYield FIXMEDOCS
	PMNoYield = 0x0002
)

// WindowProcCallback FIXMEDOCS
type WindowProcCallback func(hwnd syscall.Handle, msg uint32, wParam uintptr, lParam uintptr) uintptr
