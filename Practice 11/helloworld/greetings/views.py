from django.shortcuts import render

def greet(request, name=""):
    if name:
        greeting = f"Hello, {name}!"
    else:
        greeting = "Hello, Anonymous!"
    return render(request, 'greetings/greet.html', {'greeting': greeting})