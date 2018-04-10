# GoEnumerator
A personal tool in GO for my usual first enumeration steps on a target
---

I will release precompiled packages when first release, but in the mean time:

What do you need?
- Golang
- goquery

---
Build for GNU/Linux  
- make

Build for Darwin
- make darwin

Build for Windows
- make windows

Run:

```  
  ./GoEnumerator hispagatos.org
```


Todo:
- Add CVE feed https://nvd.nist.gov/vuln/data-feeds#JSON_FEED [done]
- Scan results based on banner and other fingerprinting against CVE's[done]
 - need to add sort to not scan two times for same banner
 - Just notice I also need to check for Vendor not just Product
 - Add found CVE's to file [done]
 - Add links to CVE's description online[done]
- make data enumerated available to all diff scans and checks.[working on this]
- Create report out of tmp files
- Create HTMl/PDF reports
- Add option to dirbust recursevely, not sure of this yet.
 - Make dirbust more eficient
- refacto all the durty/hacky code after I add all the wanted options
- Grab robots.txt[done]
- Grab comments[done]
- Get JavaScript 
- Add optional Whois for domain and IP
- Add optional dns lookup, check for more domains pointing to that ip
- Add optional dns enumeration, trying to dns brute force for more domains

![GoEnumerator Demo Animated Gif](https://github.com/ReK2Fernandez/GoEnumerator/blob/master/demo-goenumerator.gif)
