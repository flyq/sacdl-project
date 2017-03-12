var EventEmitter = require('events').EventEmitter;
var ee = new EventEmitter();

ee.on('some_events',function(foo,bar){
    console.log("第1个监听事件，参数foo=" + foo + ",bar=" + bar);
});

var isSuccess = ee.emit('some_events','Wilson','Zhong');

ee.on('some_events',function(foo,bar){
    console.log("第2个监听事件，参数foo=" + foo + ",bar="+bar);
});

ee.emit('some_events','zhong','wei');

var isSuccess2 = ee.emit('other_events','Wilson','Zhong');

console.log(isSuccess);
console.log(isSuccess2);
