package main

const (
	tpl = `	
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<html lang="en">
<head>
  <script type="text/javascript" src="http://cdn.hcharts.cn/jquery/jquery-1.8.3.min.js"></script>
  <script type="text/javascript" src="http://cdn.hcharts.cn/highcharts/highcharts.js"></script>
  <script type="text/javascript" src="http://img.hcharts.cn/highcharts/highcharts.js"></script>
  <script>
    $(function () {
    $('#insert').highcharts({
        chart: {
            type: 'spline'
        },
        title: {
            text: '{{.Headtitle}}'
        },/*
        subtitle: {
            text: 'bench from make',
            x: -20
        },*/
        xAxis: {
            categories:
			{{.Xdata}}
        },
        yAxis: {
            title: {
                text: '{{.Ytitle}}'
            },
            labels: {
                formatter: function() {
                    return this.value 
                }
            }
        }, 
        credits: {   
            text: 'Juanpi_db',
            href: 'http://www.dbauto.org'  
        }, 
        tooltip: {
            crosshairs: true,
            shared: true,
            valueSuffix: ''
        },
        plotOptions: {
            spline: {
                marker: {
                    radius: 2,
                    lineColor: '#234',
                    lineWidth: 1
                }
            }
        },
        series: [{
            name: '{{.Xtitle}}',
            data: {{.Ydata}}
	 		}
		]          
    });             
});     

 $(function () {
    $('#cpuinfo').highcharts({
        chart: {
            type: 'spline'
        },
        title: {
            text: '{{.CpuHeadtitle}}'
        },/*
        subtitle: {
            text: 'bench from make',
            x: -20
        },*/
        xAxis: {
            categories:
			{{.CpuXdata}}
        },
        yAxis: {
            title: {
                text: '{{.CpuYtitle}}'
            },
            labels: {
                formatter: function() {
                    return this.value 
                }
            }
        }, 
        credits: {   
            text: 'Juanpi_db',
            href: 'http://www.dbauto.org'  
        }, 
        tooltip: {
            crosshairs: true,
            shared: true,
            valueSuffix: ''
        },
        plotOptions: {
            spline: {
                marker: {
                    radius: 2,
                    lineColor: '#234',
                    lineWidth: 1
                }
            }
        },
        series: [
			 {{  range $index,$value := .CpuYdata}}
			 	 { name: "cpu{{$index}}" ,
				   data: {{$value}}
				  },
			{{end}}
		]          
    });             
});     

 $(function () {
    $('#meminfo').highcharts({
        chart: {
            type: 'spline'
        },
        title: {
            text: '{{.MemHeadtitle}}'
        },/*
        subtitle: {
            text: 'bench from  make',
            x: -20
        },*/
        xAxis: {
            categories:
			{{.MemXdata}}
        },
        yAxis: {
            title: {
                text: '{{.MemYtitle}}'
            },
            labels: {
                formatter: function() {
                    return this.value 
                }
            }
        }, 
        credits: {   
            text: 'Juanpi_db',
            href: 'http://www.dbauto.org'  
        }, 
        tooltip: {
            crosshairs: true,
            shared: true,
            valueSuffix: ''
        },
        plotOptions: {
            spline: {
                marker: {
                    radius: 2,
                    lineColor: '#234',
                    lineWidth: 1
                }
            }
        },
        series: [
           		 {{  range $index,$value := .MemYdata}}
					
			 	 { name: "mem{{$index}}" ,
				   data: {{$value}}
				  },
			{{end}}
	 		
		]          
    });             
});     

 </script>            
</head>
<p>sysyben</p>
<body>
 	<div id="insert"></div>
 	<div id="cpuinfo"></div>
  	<div id="meminfo"></div>
</body>
</html>

	`
)
