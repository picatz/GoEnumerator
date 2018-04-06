# GoEnumerator
A tool wrote in GO for very fast pasive enumeration of a target

---
Build for GNU/Linux  
- make

Build for Darwin
- make darwin

Build for Windows
- make windows

Run:  
- GoEnumerator hispagatos.org
- do not add http/https in front.


Todo:
- Add CVE feed https://nvd.nist.gov/vuln/data-feeds#JSON_FEED [done]
- Scan results based on banner and other fingerprinting against CVE's[done]
 - need to add sort to not scan two times for same banner
 - Just notice I also need to check for Vendor not just Product
- make data enumerated available to all diff scans and checks.[working on this]
- Create report out of tmp files
- Create HTMl/PDF reports
- Add option to dirbust recursevely, not sure of this yet.
 - Make dirbust more eficient
- refacto all the durty/hacky code after I add all the wanted options
