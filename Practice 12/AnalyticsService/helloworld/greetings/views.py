from django.shortcuts import render
from django.http import HttpResponse
import matplotlib.pyplot as plt
import io
from .models import Analytics
from django.views.decorators.csrf import csrf_exempt
import json
from django.http import JsonResponse
from django.conf import settings

def greet(request, name=""):
    if name:
        greeting = f"Hello, {name}!"
    else:
        greeting = "Hello, Anonymous!"
    return render(request, 'greetings/greet.html', {'greeting': greeting})
@csrf_exempt
def makeDiagram(request):

    # decode body
    body_unicode = request.body.decode('utf-8')
    body = json.loads(body_unicode)
    dataMassive = body["massive"]
    if dataMassive is None:
        return JsonResponse("Unknown data")
    
    productType = dataMassive[0]['productType']

    try:
        value = Analytics.objects.get(count=len(dataMassive), product_type = productType)
    except Analytics.DoesNotExist:
        value = None
    if value is not None:
        return HttpResponse(value.image, content_type="image/png")


    


    # generate new diagram

    # get prices
    prices = list(map(lambda s: s['price'], dataMassive))
    # Определяем границы для группировки продуктов по цене
    price_bins = [0, 5000, 15000]

    # Группируем продукты по цене
    price_groups = plt.hist(prices, bins=price_bins)


    # Настройки диаграммы
    plt.xlabel('Price Groups')
    plt.ylabel('Number of Products')
    plt.title('Distribution of Products by Price')

    # Создаем подписи для каждой группы
    labels = ['Low Price', 'Medium Price', 'High Price']
    print(price_bins)
    plt.xticks(price_bins, labels)


    # Загружаем Байты изображения в буффер
    buffer = io.BytesIO()
    plt.savefig(buffer, format='png')



    new_entity = Analytics(count = len(prices), product_type = productType, image=buffer.getvalue())
    new_entity.save()
    # save in postgres

    return HttpResponse(buffer.getvalue(), content_type="image/png")
