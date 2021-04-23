
# Idea
Tool to investigate "possible" caching issues in the MassIVE USI api

### What does it do
1. Spawn 50*2 (can be defined via CLI flag `-numberRequests`) [lightweight threads](https://en.wikipedia.org/wiki/Go_(programming_language)#Concurrency:_goroutines_and_channels)
2. request either `mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2` or `mzspec:PXD000394:20130504_EXQ3_MiBa_SA_Fib-2:scan:4234:SGVSRKPAPG/2` 
3. Compare the request against the response.USI

The problem might be hard to track down in other languages as you have to create highly async fast requests (simple bash/curl scripts won't help).

## Usage
```
Usage of testMassIVE:
  -numberRequests int
    	number of requests to two different endpoints each (default 50)
```

## Example
The bug is not 100 percent reproducible and you have to execute the command sometimes multiple times.
```bash
tschmidt@Tobiass-MacBook-Pro massIVE % ./builds/testMassIVE_mac
50 request against mzspec:PXD000394:20130504_EXQ3_MiBa_SA_Fib-2:scan:4234:SGVSRKPAPG/2 and mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2, respectivly

0 wrong responses%
```

This is an example where it breaks. It requests `mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2` and receives the USI `mzspec:PXD000394:ccms_peak/20130504_EXQ3_MiBa_SA_Fib-2.mzXML:scan:4234:SGVSRKPAPG/2`
```bash
tschmidt@Tobiass-MacBook-Pro massIVE % ./builds/testMassIVE_mac
50 request against mzspec:PXD000394:20130504_EXQ3_MiBa_SA_Fib-2:scan:4234:SGVSRKPAPG/2 and mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2, respectivly

request mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2
response mzspec:PXD000394:ccms_peak/20130504_EXQ3_MiBa_SA_Fib-2.mzXML:scan:4234:SGVSRKPAPG/2

request mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2
response mzspec:PXD000394:ccms_peak/20130504_EXQ3_MiBa_SA_Fib-2.mzXML:scan:4234:SGVSRKPAPG/2

request mzspec:PXD000561:Adult_Frontalcortex_bRP_Elite_85_f09:scan:17555:VLHPLEGAVVIIFK/2
response mzspec:PXD000394:ccms_peak/20130504_EXQ3_MiBa_SA_Fib-2.mzXML:scan:4234:SGVSRKPAPG/2

3 wrong responses%
```

Both results hint towards the fact that either
- Web request caching (e.g. Varnish) on IP address is done (without evaluating the exact URL) -> possible security hole
- Database requests are cached -> same here