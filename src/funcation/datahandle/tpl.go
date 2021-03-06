package datahandle

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
            text: 'bench from  ds_group@juanpi.com',
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
            type: 'spline',
			zoomType: 'x'
        },
        title: {
            text: '{{.CpuHeadtitle}}'
        },
        subtitle: {
            text: 'bench from  ds_group@juanpi.com',
            x: -20
        },
        xAxis: {
             type: 'datetime',
            labels: {
                overflow: 'justify'
            }
        },
        yAxis: {
            title: {
                text: '{{.CpuYtitle}}'
            },
			min: 0,
            minorGridLineWidth: 0,
            gridLineWidth: 0,
            alternateGridColor: null
        }, 
        credits: {   
            text: 'Juanpi_db',
            href: 'http://www.dbauto.org'  
        }, 
       tooltip: {
            valueSuffix: 'cpu idle'
        },
          plotOptions: {
            spline: {
                lineWidth: 2,
                states: {
                    hover: {
                        lineWidth: 1
                    }
                },
                marker: {
                    enabled: false
                },
                pointInterval: {{.Interval}}, // one 秒 1000 = 1s
                pointStart: Date.UTC(
					        {{range $k,$v:= .StartTime}}
                                   {{$v}}
                            {{ if lt  $k 5 }}
                                ,
                            {{end}}
						{{end}}
				)
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
            type: 'spline',
			zoomType: 'x'
        },
        title: {
            text: '{{.MemHeadtitle}}'
        },/*
        subtitle: {
            text: 'bench from  ds_group@juanpi.com',
            x: -20
        },*/
        xAxis: {
             type: 'datetime',
            labels: {
                overflow: 'justify'
            }
        },
        yAxis: {
            title: {
                text: '{{.MemYtitle}}'
            },
            labels: {
                formatter: function() {
                    return this.value 
                }
            },
			min: 0,
            minorGridLineWidth: 0,
            gridLineWidth: 0,
            alternateGridColor: null
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
        tooltip: {
            valueSuffix: 'cpu idle'
        },
          plotOptions: {
            spline: {
                lineWidth: 2,
                states: {
                    hover: {
                        lineWidth: 1
                    }
                },
                marker: {
                    enabled: false
                },
                pointInterval: {{.Interval}}, // one 秒 1000 = 1s
                pointStart: Date.UTC(
					        {{range $k,$v:= .StartTime}}
                            	{{$v}}
                           	 	{{ if lt  $k 5 }}
                               		 ,
                           		 {{end}}
							{{end}}
				)
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
<body>
 	<div id="insert"></div>
 	<div id="cpuinfo"></div>
  	<div id="meminfo"></div>
</body>
</html>
`
)
