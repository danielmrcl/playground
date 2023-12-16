from .models import Question
from django.http import HttpRequest, HttpResponse, Http404
from django.shortcuts import get_object_or_404, render

def index(request: HttpRequest):
    questions = Question.objects.all()
    return render(request, "polls/index.html", { "questions": questions })

def details(request, question_id):
    question = get_object_or_404(Question, id=question_id)
    return render(request, "polls/details.html", { "question": question })

def results(request, question_id):
    return HttpResponse("There are the results for question %s." % question_id)

def votes(request, question_id):
    return HttpResponse("This question %s have received _ votes." % question_id)

