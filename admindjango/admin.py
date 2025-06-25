from django.contrib import admin
from .models import Users, Courses, Materials, Assignments, StudentUploads  

admin.site.register(Users)
admin.site.register(Courses)
admin.site.register(Materials)
admin.site.register(Assignments)
admin.site.register(StudentUploads)
