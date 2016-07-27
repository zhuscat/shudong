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
                    alert("delete success");
                    //all delete is success,this can be execute
                    //$("#tr_"+comment_id).remove();
                    $(this).css("background","#dedada");
                }
                }
            });
        }
};

function gag(user_id){
 confirm_=confirm('This action will gag! Are you sure?');
        if(confirm_){
            $.ajax({
                type:"POST",
                url:'../static/data/banCommitComment.json',
                data:{
                    id:user_id
                },
                success:function(data){
                    if(data.success){
                    alert("gag success");
                    //all delete is success,this can be execute
                    //$("#tr_"+comment_id).remove();
                    $(this).css("background","#5cb85c");
                    $(this).innerHTML="解除禁言";
                }
                }
            });
        }
};

function removegag(user_id){
 confirm_=confirm('This action will remove the gag! Are you sure?');
        if(confirm_){
            $.ajax({
                type:"POST",
                url:'../static/data/removegag.json',
                data:{
                    id:user_id
                },
                success:function(data){
                    if(data.success){
                    alert("remove the gag success");
                    //all delete is success,this can be execute
                    //$("#tr_"+comment_id).remove();
                    $(this).css("background","#d9534f");
                    $(this).innerHTML="禁言";
                }
                }
            });
        }
};

function undercarriage(book_id){
 confirm_=confirm('This action will remove the book! Are you sure?');
        if(confirm_){
            $.ajax({
                type:"POST",
                url:'../static/data/undercarriage.json',
                data:{
                    id:user_id
                },
                success:function(data){
                    if(data.success){
                    alert("remove the book success");
                    //all delete is success,this can be execute
                    //$("#tr_"+comment_id).remove();
                    $(this).css("background","#d9534f");
                    $(this).innerHTML="已下架";
                }
                }
            });
        }
};


function deletebook(book_id){
 confirm_=confirm('This action will delete the book! Are you sure?');
        if(confirm_){
            $.ajax({
                type:"POST",
                url:'../static/data/deletebook.json',
                data:{
                    id:book_id
                },
                success:function(data){
                    if(data.success){
                    alert("delete the book success");
                    //all delete is success,this can be execute
                    //$("#tr_"+comment_id).remove();
                    $(this).css("background","#fafafa");
                    $(this).innerHTML="已删除";
                }
                }
            });
        }
};