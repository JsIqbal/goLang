একটি Go ফাইল সাধারনত নিম্নলিখিত অংশগুলো  নিয়ে গঠিত – 

প্যাকেজ ডিক্লেয়ারেশন
প্যাকেজ ইমপোর্ট 
ফাংশন
স্টেটমেন্ট এবং এক্সপ্রেশন

reserved keyword:
break     default      func    interface  select
case      defer        go      map        struct
chan      else         goto    package    switch
const     fallthrough  if      range      type
continue  for          import  return     var

Project with external package needs the mod file:
Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module
        'go mod init example.com/m/v2' to initialize a v2 module