function delete_comment(comment_id){
 confirm_=confirm('This action will delete current order! Are you sure?');
        if(confirm_){
            $.ajax({
                type:"POST",
                url:'../static/data/delete_comment.json',
                data:{
                	id:comment_id
                },
                success:function(data){
                    if(data.success){
                    alert("test order");
                    //all delete is success,this can be execute
                    //$("#tr_"+comment_id).remove();
                    $(this).css("background","#dedada");
                }
                }
            });
        }
    };
