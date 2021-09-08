proc_list = [ "ps1",
      "ldf",
      "lock",
      "theme",
      "msi",
      "sys",
      "wpx",
      "cpl",
      "adv",
      "msc",
      "scr",
      "bat",
      "key",
      "ico",
      "dll",
      "hta",
      "deskthemepack",
      "nomedia",
      "msu",
      "rtp",
      "msp",
      "idx",
      "ani",
      "386",
      "diagcfg",
      "bin",
      "mod",
      "ics",
      "com",
      "hlp",
      "spl",
      "nls",
      "cab",
      "exe",
      "diagpkg",
      "icl",
      "ocx",
      "rom",
      "prf",
      "themepack",
      "msstyles",
      "lnk",
      "icns",
      "mpa",
      "drv",
      "cur",
      "diagcab",
      "cmd",
      "shs"]

proc_list2 = ["themepack","nls","diagpkg","msi","lnk","exe","cab","scr","bat","drv","rtp","msp","prf","msc","ico","key","ocx","diagcab","diagcfg","pdb","wpx","hlp","icns","rom","dll","msstyles","mod","ps1","ics","hta","bin","cmd","ani","386","lock","cur","idx","sys","com","deskthemepack","shs","ldf","theme","mpa","nomedia","spl","cpl","adv","icl","msu"]

result = set()

for each in proc_list:
    result.add(each)
    
for each in proc_list2:
    result.add(each)
    
print(result)