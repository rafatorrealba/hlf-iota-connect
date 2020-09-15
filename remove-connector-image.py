import subprocess

images = subprocess.run(['docker', 'images'], stdout=subprocess.PIPE)
connector = images.stdout.decode('utf-8')[179:350]
print(connector)

id = connector[134:146]
print(id)

subprocess.run(['docker', 'rmi', '-f', id])
