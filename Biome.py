#!/usr/local/bin/python3

# MIT License

# Copyright (c) 2017 Clayton Dorsey
# https://cydy.io @claytondorsey 

import numpy as np
from EmojiGrid import *
import random
from chars import *

def getSectionCords(inputCanvas, sectNum):
    return np.argwhere(inputCanvas == sectNum)

def splitNum(parts, total):
    elements = (np.arange(parts) == np.random.randint(0, parts) for i in range(total))
    return np.sum(elements, axis=0).tolist()

class Biome:

    def replaceSection(self, grid_input, cur, new):
        cords = np.argwhere(grid_input == cur)
        grid_input[cords[:, 0], cords[:, 1]] = new
    
    def crops(self):
        #if random.randint(1, 3) == 1:
        if 1 == 1:
            print('crops: whole')
            crop1 = random.choice(PLANTS_CROPS)
            print('crop:', crop1)

            self.replaceSection(self.sectionCanvas, self.sectNum, crop1)
        pass

    def animals(self):
        placementCords = self.cords
        numberOfCords = placementCords.shape[0] #returns number of cordinates
        fillPercentage = random.uniform(.2, .6)
        numCordsToFill = int(round(numberOfCords * fillPercentage))
        print('filling', numberOfCords, fillPercentage, numCordsToFill)
        placementCords = placementCords[np.random.choice(placementCords.shape[0], numCordsToFill, replace=False), :]

        self.replaceSection(self.sectionCanvas, self.sectNum, ':seedling:') #replace all of section with em-space

        if random.randint(0, 1) == 1:
            print('animals: livestock')
            animalSet = ANIMALS_FARM
        else:
            print('animals: birds')
            animalSet = ANIMALS_BIRDS

        if random.randint(0, 3) == 0: # 3/4 chance one animal biome
            print('animals #: 1')
            animal1 = random.choice(animalSet)
            print('animal1:', animal1)
            for x in placementCords:
                self.sectionCanvas[x[0]][x[1]] = animal1
                pass

        else:
            if 5 > numCordsToFill:
                animalMax = numCordsToFill
            else:
                animalMax = 5
            
            animalNum = random.randint(2, animalMax)
            print('animals #:', animalNum)
            animalChoices = random.sample(animalSet, animalNum)
            animalDistribution = splitNum(animalNum, numCordsToFill)
            #print(animalChoices)
            #print(animalDistribution)
            #print(placementCords)

            z = 0
            for y, animal_output in zip(animalDistribution, animalChoices):
                
                for x in range(y):
                    #print('running ' + str(x) + ' out of', str(y), 'times')
                    #print('placing', animalChoices[animalDistribution.index(y)], 'at',  str(placementCords[z+x][0]), str(placementCords[z+x][1]))
                    self.sectionCanvas[placementCords[z+x][0]][placementCords[z+x][1]] = animal_output
                z += y

        pass

    def field(self):
        placementCords = self.cords
        numberOfCords = placementCords.shape[0] #returns number of cordinates
        fillPercentage = random.uniform(.2, .8)
        numCordsToFill = int(round(numberOfCords * fillPercentage))
        print('filling', numberOfCords, fillPercentage, numCordsToFill)
        placementCords = placementCords[np.random.choice(placementCords.shape[0], numCordsToFill, replace=False), :]

        self.replaceSection(self.sectionCanvas, self.sectNum, ':seedling:') #replace all of section with em-space

        if random.randint(0, 1) == 1:
            print('plants: field')
            plantSet = PLANTS_FIELD
        else:
            print('plants: flowers')
            plantSet = PLANTS_FLOWERS

        if random.randint(0, 3) == 0: # 3/4 chance one plant biome
            print('plants #: 1')
            plant1 = random.choice(plantSet)
            print('plant1:', plant1)
            for x in placementCords:
                self.sectionCanvas[x[0]][x[1]] = plant1
                pass

        else:
            if 5 > numCordsToFill:
                plantMax = numCordsToFill
            else:
                plantMax = 5
            
            plantNum = random.randint(2, plantMax)
            print('plants #:', plantNum)
            plantChoices = random.sample(plantSet, plantNum)
            plantDistribution = splitNum(plantNum, numCordsToFill)
            #print(plantChoices)
            #print(plantDistribution)
            #print(placementCords)

            z = 0
            for y, plant_output in zip(plantDistribution, plantChoices):
                
                for x in range(y):
                    #print('running ' + str(x) + ' out of', str(y), 'times')
                    #print('placing', plantChoices[plantDistribution.index(y)], 'at',  str(placementCords[z+x][0]), str(placementCords[z+x][1]))
                    self.sectionCanvas[placementCords[z+x][0]][placementCords[z+x][1]] = plant_output
                z += y

        pass

    def pond(self):
        placementCords = self.cords
        numberOfCords = placementCords.shape[0] #returns number of cordinates
        fillPercentage = random.uniform(.1, .5)
        numCordsToFill = int(round(numberOfCords * fillPercentage))
        print('filling', numberOfCords, fillPercentage, numCordsToFill)
        placementCords = placementCords[np.random.choice(placementCords.shape[0], numCordsToFill, replace=False), :]

        self.replaceSection(self.sectionCanvas, self.sectNum, ':water_wave:') #replace all of section with em-space

        pawnSet = POND

        if random.randint(0, 3) == 0: # 3/4 chance one pawn biome
            print('pawns #: 1')
            pawn1 = random.choice(pawnSet)
            print('pawn1:', pawn1)
            for x in placementCords:
                self.sectionCanvas[x[0]][x[1]] = pawn1
                pass

        else:
            if 5 > numCordsToFill:
                pawnMax = numCordsToFill
            else:
                pawnMax = 5
            
            pawnNum = random.randint(2, pawnMax)
            print('pawns #:', pawnNum)
            pawnChoices = random.sample(pawnSet, pawnNum)
            pawnDistribution = splitNum(pawnNum, numCordsToFill)
            #print(pawnChoices)
            #print(pawnDistribution)
            #print(placementCords)

            z = 0
            for y, pawn_output in zip(pawnDistribution, pawnChoices):
                
                for x in range(y):
                    #print('running ' + str(x) + ' out of', str(y), 'times')
                    #print('placing', pawnChoices[pawnDistribution.index(y)], 'at',  str(placementCords[z+x][0]), str(placementCords[z+x][1]))
                    self.sectionCanvas[placementCords[z+x][0]][placementCords[z+x][1]] = pawn_output
                z += y

        pass

    def barrier(self):
            #if random.randint(1, 3) == 1:
            if 1 == 1:
                print('barrier: whole')
                barrier1 = random.choice(BARRIERS)
                print('barrier:', barrier1)

                self.replaceSection(self.sectionCanvas, self.sectNum, barrier1)
            pass
    
    def __init__(self, farm_instance, sectionCanvas, sectionNum):
        #super().__init__()
        self.farm = farm_instance
        self.sectionCanvas = sectionCanvas
        self.sectNum = sectionNum
        self.cords = getSectionCords(self.sectionCanvas, self.sectNum)

        print('\nsection number:', str(self.sectNum))

        self.available_biome_types = {0 : self.crops,
                                1 : self.animals,
                                2 : self.field,
                                3 : self.pond}

        if int(self.sectNum) % 2 == 0:
            #print('biome type: BARRIER\n')
            pass
        else:
            self.BIOME_TYPE = random.choice(list(self.available_biome_types.keys()))
            #print('biome type:', self.BIOME_TYPE)
        
        #dimensions = abs(self.cords[-1] - self.cords[0])
	    #sml = np.ones(dimensions + 1)
        #self.canvas = np.full((dimensions + 1), self.sectNum, dtype='U64')
        self.build()
        self.replace()

    def build(self):
        if int(self.sectNum) % 2 == 0:
            print('biome type: BARRIER\n')
            self.barrier()
        else:
            #self.BIOME_TYPE = random.choice(list(self.available_biome_types.keys()))
            #print('biome type:', self.BIOME_TYPE)
            self.available_biome_types[self.BIOME_TYPE]()
        
    def replace(self):
        np.place(self.farm.canvas_strings, self.farm.canvas_strings == self.sectNum, self.sectionCanvas)

    def __str__(self):
        return '\n'.join([''.join([str(elem) for elem in row]) for row in self.sectionCanvas]) + '\n'

    