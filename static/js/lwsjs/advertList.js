  

//判断是否有选中的checkbox
function Select()
{
    var array = "";//变量用来存储选中checkbox的Value

    if ($("input[name='check']:checked").size() == 0)
    {
          $.dialog({
          title:"提示",
          content: '尚未选择任何对象',
          icon:"alert.gif",

          ok: function(){
      
          return true ;
                        },
                    });

    }
    else
    {
        $("input[name='check']:checked").each(function ()
        {
            //alert($(this).attr("value"));
            if (array == "")
            {
                array = $(this).attr("value");
            }
            else
            {
                array = array + "," + ($(this).attr("value"));
            }
        }
        );
        return array;
    }
}



      function Del(en,href)
        {

       
          if ($(en).attr("class")=="del") 
          {
          	var data=Select();

             if (data) 
             {

                $.dialog.confirm('你确定要删除这些数据吗？', function()
                         {
                         
                            window.location.href=href+"?str="+data;

                      

                         }, function()
                         {
   
                         });
 
             };

                 
          }else
          {
           $.dialog.confirm('你确定要删除这条数据吗？', function(){
               var va= $(en).siblings("input:hidden").attr("value");
              
                $(en).attr("href",href);
                  
                 window.location.href=$(en).attr("href");
                                  //this.title('3秒后自动关闭').time(3);
            }, function(){
   
           });
          }
        }





      //全选多选
      function CheckAll(en) {
          
          var text=$(en).find("span"). html();
          if (text=="全选")
          {

             $(".checkbox").each(function() {
                     this.checked = true;
                    });
              
              $(en).find("span"). html("全不选");
          }
          else
          {


                    $(".checkbox").each(function() {
                     this.checked = false;
                    });
              
              $(en).find("span"). html("全选");
          }

                
      }