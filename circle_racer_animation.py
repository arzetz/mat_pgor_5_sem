import pygame
import math
import time
import random

window_size = 600
radius = 200 
circle_center = (window_size // 2, window_size // 2)
speed = 2 * math.pi 
fps = 60

pygame.init()
screen = pygame.display.set_mode((window_size, window_size))
pygame.display.set_caption("Гонщики на окружности")
clock = pygame.time.Clock()

white = (255, 255, 255)
blue = (0, 0, 255)
red = (255, 0, 0)
green = (0, 255, 0)
some = (177, 65, 200)
some2 = (0, 255, 255)
some3 = (255, 0, 255)

class Racer:
    def __init__(self, color, speed_offset, name):
        self.color = color
        self.speed_offset = speed_offset
        self.laps = 0
        self.prev_angle = 0
        self.name = name
        self.last_speed_change_time = time.time()

    def update_position(self, time_elapsed):
        angle = (speed * time_elapsed) * self.speed_offset / 5
        x = circle_center[0] + int(radius * math.cos(angle))
        y = circle_center[1] + int(radius * math.sin(angle))

        current_angle = angle % (2 * math.pi)
        if self.prev_angle > current_angle:
            self.laps += 1

        self.prev_angle = current_angle
        return x, y
    
    def update_speed(self):
        if time.time() - self.last_speed_change_time > 2:
            self.speed_offset = math.pi - random.uniform(-0.05, 0.05)
            self.last_speed_change_time = time.time()

racers = [
    Racer(red, math.pi - random.uniform(-0.3, 0.3), "Ландо"),
    Racer(some, math.pi - random.uniform(-0.3, 0.3), "Льюис"),
    Racer(some2, math.pi - random.uniform(-0.3, 0.3), "Макс"),
    Racer(some3, math.pi - random.uniform(-0.3, 0.3), "Эстебан"),
]
start_time = time.time()

running = True
show_text = False
final_text = "Все гонщики завершили 3 круга"

while running:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False
    time_elapsed = time.time() - start_time

    if not show_text and all(r.laps >= 5 for r in racers):
        show_text = True
        lap_data = [(r, r.laps) for r in racers]
        lap_data.sort(key=lambda x: x[1], reverse=True)
        screen.fill(white)
        font = pygame.font.Font(None, 36)

        text_y = window_size // 2 - len(racers) * 20
        final_text = "Результаты гонщиков:"
        title = font.render(final_text, True, green)
        screen.blit(title, (window_size // 2 - title.get_width() // 2, text_y))

        for i, (racer, laps) in enumerate(lap_data):
            result_text = f'{racer.name}: {laps} кругов'
            text_surface = font.render(result_text, True, racer.color)
            screen.blit(text_surface, (window_size // 2 - text_surface.get_width() // 2, text_y + (i + 1) * 40))

        pygame.display.flip()
        time.sleep(10)
        break

    if not show_text:
        screen.fill(white)
        pygame.draw.circle(screen, blue, circle_center, radius, 1)

        for racer in racers:
            racer.update_speed()
            x, y = racer.update_position(time_elapsed)
            pygame.draw.circle(screen, racer.color, (x, y), 10)
            current_position = [(racer, racer.laps) for racer in racers]
            current_position.sort(key=lambda x: x[1], reverse=True)
        font = pygame.font.Font(None, 36)
        for i, (racer, laps) in enumerate(current_position):
            text = font.render(f'{racer.name}: {laps} кругов', True, racer.color)
            screen.blit(text, (20, 20 + i * 40))

        pygame.display.flip()

        clock.tick(fps)

pygame.quit()
