from django.db import models

class Question(models.Model):
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField(auto_now_add=True)

class Choice(models.Model):
    question = models.ForeignKey(to=Question, on_delete=models.SET_NULL, null=True)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)
