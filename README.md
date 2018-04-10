# GoEnumerator
A personal tool in GO for my usual first enumeration steps on a target
---

I will release precompiled packages when first release, but in the mean time:

What do you need?
- Go
- Goquery - go get github.com/PuerkitoBio/goquery

---
Build for GNU/Linux  
- make

Build for Darwin
- make darwin

Build for Windows
- make windows


Configure:
- Edit conf.json and change it to your personal dictionary for web enumeration and password dictionary


Run:

```  
  ./GoEnumerator hispagatos.org
```


# To do
- [x] Add CVE feed https://nvd.nist.gov/vuln/data-feeds#JSON_FEED.
- [x] Scan results based on banner and other fingerprinting against CVE's.
 - [ ] need to add sort to not scan two times for same banner.
 - [ ] Just notice I also need to check for Vendor not just Product.
 - [x] Add found CVE's to file.
 - [x] Add links to CVE's description online.
- [x] Make data enumerated available to all diff scans and checks.
- [ ] Create report out of tmp files.
- [ ] Create HTMl/PDF reports.
- [ ] Add option to dirbust recursevely, not sure of this yet.
 - [ ] Make dirbust more eficient.
- [ ] Refactor all the durty/hacky code after I add all the wanted options.
- [x] Grab robots.txt.
- [x] Grab comments.
- [ ]  Get JavaScript. 
- [ ] Add optional Whois for domain and IP.
- [ ] Add optional dns lookup, check for more domains pointing to that ip.
- [ ] Add optional dns enumeration, trying to dns brute force for more domains.


![GoEnumerator Demo Animated Gif](https://github.com/ReK2Fernandez/GoEnumerator/blob/master/demo-goenumerator.gif)
