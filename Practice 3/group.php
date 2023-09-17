<?php
header('Content-Type: application/json');
$requestBody = file_get_contents("php://input"); // получаем тело запроса в виде строки
$data = json_decode($requestBody, true);
$mysqli = new mysqli("db", "user", "password", "appDB");
$result;

// POST
if($_SERVER["REQUEST_METHOD"] == "POST"){
    if(isset($data['name']))
    {
    $result = $mysqli->query("INSERT INTO `group` (`name`) values ('{$data['name']}');");
    header("HTTP/1.1 201 Created");

    }
    else {
        header("HTTP/1.1 400 Bad Request");

    }
}
// GET
elseif($_SERVER["REQUEST_METHOD"] == "GET"){
    if(isset($_GET["id"])) {
        $result = $mysqli->query ("SELECT * FROM `group`where ID = {$_GET['id']}");
    }
    elseif(isset($_GET["name"])) {
        $result = $mysqli->query("SELECT * FROM `group` where name = '{$_GET['name']}'");
    }
    else {
        $result = $mysqli->query("SELECT * FROM `group`");
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
elseif($_SERVER["REQUEST_METHOD"] == "PUT"){
    if(isset($data))
    {
     $result = $mysqli->query("UPDATE `group` set `name`='{$data['name']}' where `ID` = {$data['ID']};");
     header("HTTP/1.1 200 OK");
   
   }
   else {
     header("HTTP/1.1 400 Bad Request");
   }
}
// DELETE
elseif($_SERVER["REQUEST_METHOD"] == "DELETE") {
    if(isset($data))
    {
   
   $result = $mysqli->query("DELETE FROM `group` where `ID`={$data['ID']};");
   header("HTTP/1.1 200 OK");
   
    }
    else {
       header("HTTP/1.1 400 Bad Request");
   
    }
}
else {
    header("HTTP/1.1 405 Method Not Allowed");
}
?>