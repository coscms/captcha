(function(){
    var pathEndSepRegex=new RegExp('/$'),pathPlaceholderRegex=new RegExp('\{(driver|type)\}','g');
    var defaults={driver:'click',type:'basic',api:'/captcha/',success:null,error:null,networdError:null,complete:null};
    var captcha=function(container,settings) {
        this.container=$(container);
        this.options=$.extend({},defaults,settings||{});
        this.captcha=null;
        this.init();
    };
    captcha.prototype.init=function(){
        var _this=this,apiURLMaker;
        if(pathPlaceholderRegex.test(this.options.api)){
            apiURLMaker=function(driver,type){
                return _this.options.api.replace(pathPlaceholderRegex,function(v){
                    if(v=='{driver}') return driver;
                    return type;
                })
            }
        }else{
            if(!pathEndSepRegex.test(this.options.api)) this.options.api+='/';
            apiURLMaker=function(driver,type){
                return _this.options.api+driver+'/'+type;
            }
        }
        var api=apiURLMaker(this.options.driver,this.options.type);
        var opts={dataApi:api,verifyApi:api,container:container,success:this.options.success,error:this.options.error,networdError:this.options.networkError,complete:this.options.complete};
        switch(this.options.driver){
            case 'click':
                this.captcha=this.options.type=='shape'?CaptchaClickShape(opts):CaptchaClickBasic(opts);break;
            case 'rotate':
                this.captcha=CaptchaRotate(opts);break;
            case 'slide':
                this.captcha=this.options.type=='region'?CaptchaSlideRegion(opts):CaptchaSlideBasic(opts);break;
            default:
                opts.dataApi=apiURLMaker('click','basic');
                opts.verifyApi=opts.dataApi;
                this.captcha=CaptchaClickBasic(opts);break;
        }
    };
    $.fn.captcha=function(settings) {
        return this.each(function() {
            return captcha(this,settings);
        });
    };
    $.fn.captcha.Constructor = captcha;
})();