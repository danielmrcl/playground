from django.http import HttpRequest, HttpResponse
from django.shortcuts import render

def index(request: HttpRequest):
    print(request.headers)
    return HttpResponse("Index")
