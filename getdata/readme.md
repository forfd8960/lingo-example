## get salary data

```sh
./getdata script.gd
cityName:  beijing
result:  beijing
{"country":"CN","name":"beijing","val":1600}
```

```sh
 > (getsalary "beijing")
cityName:  beijing
result:  beijing
{"country":"CN","name":"beijing","val":1600}
 > (getsalary "lanzhou")
cityName:  lanzhou
result:  lanzhou
{"country":"CN","name":"lanzhou","val":500}
 > (getsalary "chengdou")
cityName:  chengdou
city chengdou salary data not found
 > ^C
```