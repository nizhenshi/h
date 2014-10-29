// JavaScript Document
/*
*工具提示控件 v1.0
*制作：小凯工作室--地瓜
*日期：2010-9-16
*QQ：648173085
*
*使用：
*    1.简易方法在元素上直接设置 tooltip="在提示的内容" 还可以同时在class里设置参数，如：class="tooltip-100-left" //[tootip-宽-对齐方式]  例：<a class="tooltip-200-center" tooltip="提示内容" herf='#'>tooltip</a>
*    2.程序使用 如：$(".toolT").toolTip({"width":"auto","text-align":"center"}); //参数是设定提示的样式
*
*升级记录：
*     无
*
*/
 
$.fn.extend({
	"toolTip":function(css){
		var cssop=$.extend({
           "border":"3px solid #CECECE",
		   "width":220,
		   "position":"absolute",
		   "left":0,
		   "top":-500,
		   "background-color":'#fcfcfc',
		   "cursor":"default",
		   "padding":"3px",
		   "z-index":200
		}, css || {});
		
		var _self=$(this);
	    _self.each(function(i){
			if($(this).css("cursor")=="auto"){
				$(this).css({"cursor":cssop.cursor});
			}
			$(this).unbind("mouseover");
			$(this).bind("mouseover",function(e){			     //mouseenter 
				if(!$(this).data("title"))
				{
					var title="";
					if($(this).attr("toolTip"))
					{
						title=$(this).attr("toolTip");	
					}
					else
					{
					    title=$(this).attr("title");
					}
					var r_reg=/^\[[a-zA-Z]+\]$/;
					if(r_reg.test(title))
					{
					   switch(title)
					   {
					       case "[this]":
						     title=$("<p></p>").append($(this).clone().css({"width":"auto","height":"auto"})).html();	 
						   break;
					   }
					}
				    $(this).data("title",title);
					$(this).removeAttr("title");
				}
				
                if($(this).data("title")==""){return false;}
				e.stopPropagation(); //防止冒泡
				
				var _tool=_createToolTip();
				_tool.html($(this).data("title")); 
				
				var w_w=$(window).width();
				
				
				
				_tool.stop().animate({"opacity":1},500);
				$(this).unbind("mousemove");
				$(this).bind("mousemove",function(em){         //mousemove
					var w_h=$(window).height();
					
					/*计算left*/
					var w_sl=$(window).scrollLeft();					
					var _lefte=em.clientX+w_sl +10;//定位left					
					var _rig=w_w+w_sl-20;          //显示右侧范围		
					if(_lefte+_tool.width()>=_rig)
					{
						_lefte=em.clientX+w_sl-_tool.width()-10;
						if(_lefte<=w_sl)
						{
						   _lefte=w_sl +10;							
						}
					}
					
					/*计算top*/
					var w_st=$(window).scrollTop();
					var _tope=em.clientY+w_st+17;//top					
					var _bitt=w_h+w_st-20;//显示底侧范围
					if(_tope+_tool.height()>=_bitt)
					{
						_tope=em.clientY+w_st-_tool.height()-15;
						if(_tope<=w_st)
						{
						   _tope=w_st+10;							
						}
					}
					
					
				    _tool.css({"top":_tope,"left":_lefte});
				    
				});
				$(this).unbind("mouseleave");
				$(this).bind("mouseleave",function(em){                   //mouseleave
				     _tool.stop().animate({"opacity":0},500,function(){
					 	$(this).css({"left":-5000});						
					 }); 
				});
				
				
			});	//bind mouseenter end
			
		});//each end
	 
	 
	 function _createToolTip(){
		if($("div._ToolTip_div").size()>0)
		{
			$("div._ToolTip_div").css(cssop);		    
		}else
		{
		    $("div:visible:first").append("<div class='_ToolTip_div'>welcome<div>");
		    $("div._ToolTip_div").css(cssop);
		}
		return $("div._ToolTip_div");
	 }
	 
	}
});//end

$(function(){
	$("[tooltip]").each(function(i){
	    var reg=/\btooltip-[0-9a-z]+(-[a-zA-Z]+)*\b/i;
		var classN=$(this).get(0).className;
		
		var _wid=classN.match(reg);
		var _cls="";
		for(i in _wid)
		{
			 if(i==0)
			 {
				_cls= _wid[i]
			    break;
			 }
		}
		var _arr=_cls.split('-');
		var _w=220;
		var _ta="left";
		switch(_arr.length)
	    {
			case 2:
			    _w= parseInt(_arr[1]);
				_w=isNaN(_w)?"auto":_w; 
		    break;
			case 3:
				_w= parseInt(_arr[1]);
				_w=isNaN(_w)?"auto":_w; 
				_ta=_arr[2];
		    break;
		} 
		$(this).toolTip({"width":_w,"text-align":_ta});
	});
});

