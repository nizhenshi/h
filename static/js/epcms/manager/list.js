$(document).ready(function(){
    $("#btnDelete").click(function () {
         if ($(".checkall input:checked").size() < 1) {
            $.dialog.alert('对不起，请选中您要操作的记录！');
            return false;
        }
        $.dialog.confirm("删除管理员后不可恢复，确定要删除？", function () {
            var proId = "";
            var isTrue = true;
            var dataMessage = "";
            $(".checkall input:checked").each(function (i) {
                proId += $(this).attr("attrId") + ",";
            });
            if (proId!=""){
                proId=proId.toString().substring(0,proId.length-1)
            }
            $.ajax({
                url: "/admin/manager/delete",
                data: { id: proId },
                type: "post",
                dataType: "json",
                async: false,
                success: function (data) {
                    dataMessage = data.msg;
                    if (data.status==1) {
                        isTrue = true;
                    }
                    else {
                        isTrue = false;
                    }
                }
            });
            $.dialog.alert(dataMessage)
            if (isTrue) {
                var strClass = proId.split(',');
                for (var i = 0; i < strClass.length; i++) {
                    $("#" + strClass[i]).remove();
                }
            }
        });
    });
});
