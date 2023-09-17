<?php
header('Content-Type: application/json');
$result;
$mysqli = new mysqli("db", "user", "password", "appDB");
$requestBody = file_get_contents("php://input"); // получаем тело запроса в виде строки
$data = json_decode($requestBody, true);
// GET
if($_SERVER["REQUEST_METHOD"] == "GET"){
    if(isset($_GET["id"])) {
        $result = $mysqli->query("SELECT * FROM student where ID = {$_GET['id']}");
    }
    elseif(isset($_GET["name"])) {
        $result = $mysqli->query("SELECT * FROM student where fullname = '{$_GET['name']}'");
    }
    elseif(isset($_GET["group_id"])) {
        $result = $mysqli->query("SELECT * FROM student where group_id = {$_GET['group_id']}");
    }
    else {
        $result = $mysqli->query("SELECT * FROM student");

    }
    $rows = array();
    if($result == true) {
        foreach ($result as $row){
            $rows[] = $row;
        }
    }
    $data = json_encode($rows);
    echo $data; 
}

// PUT
elseif($_SERVER["REQUEST_METHOD"] == "PUT") {
    if(isset($data))
    {   
     $result = $mysqli->query("UPDATE `student` set `fullname`='{$data['fullname']}', `group_id`={$data['group_id']} where `ID` = {$data['ID']};");
     header("HTTP/1.1 200 OK");
   }
   else {
     header("HTTP/1.1 400 Bad Request");
   
   }
}

// POST
elseif($_SERVER["REQUEST_METHOD"] == "POST") {
    if(isset($data['fullname']) && isset($data['group_id']))
     {
      $result = $mysqli->query("INSERT INTO `student` (`fullname`, `group_id`) values ('{$data['fullname']}', {$data['group_id']});");
      header("HTTP/1.1 201 Created");
     }
     else {
      header("HTTP/1.1 400 Bad Request");
    
     }
}

// DELETE 
elseif($_SERVER["REQUEST_METHOD"] == "DELETE"){
    if(isset($data))
    {
   $result = $mysqli->query("DELETE FROM `student` where `ID`={$data['ID']};");
   header("HTTP/1.1 200 OK");
   
    }
    else {
       header("HTTP/1.1 400 OK");
   
    }
}

else {
    header("HTTP/1.1 405 Method Not Allowed");
}
?>