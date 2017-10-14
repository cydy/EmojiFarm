#!/usr/local/bin/python3

# MIT License

# Copyright (c) 2017 Clayton Dorsey
# https://cydy.io @claytondorsey 

import random
import numpy as np
import emoji
from Biome import *

class EmojiGrid:
    rows = 10  # y axis
    cols = 13  # x axis
    size = (rows, cols)
    y_length = rows - 1
    x_length = cols - 1
    edge_buffer = 3  # 3 array space buffer on side
    filler = '-'
    canvas = np.array([(1, 2, 3), (8, 9, 4), (7, 6, 5)])

    def replaceSection(self, grid_input, cur, new):
        cords = np.argwhere(grid_input == cur)
        grid_input[cords[:, 0], cords[:, 1]] = new

    # def __init__(self):
    #     print(self)

    def __str__(self):
        return '\n'.join([''.join([str(elem) for elem in row]) for row in self.canvas]) + '\n'

    def getRow(self, row_num):
        return self.canvas[row_num, :]

    def getCol(self, col_num):
        return self.canvas[:, col_num]

    def getSections(self):
        return np.unique(self.canvas)
    
    def getBiomeNums(self):
        temp = []
        for x in self.getSections():
            if x % 2 != 0 and x != 9:
                temp.append(x)

        return temp

    def getBarrierNums(self):
        temp = []
        for x in self.getSections():
            if x % 2 == 0:
                temp.append(x)

        return temp

    def getSectionCords(self, sectNum):
        return np.argwhere(self.canvas == sectNum)

    def splitGen(self):
        #split grid into sections

        # possible direction:
        # 0 = none (no line on this axis)
        # 1 = keep '2' or '4' barrier
        # 2 = keep '6' or '8' barrier
        # 3 = keep both

        dir_x = random.randint(0, 3)
        dir_y = random.randint(0, 3)

        if ((dir_x == 0) or (dir_y == 0) and (((dir_x == 1) or (dir_y == 1)) or ((dir_x == 2) or (dir_y == 2)))):
            print('no split')
            # maybe add second chance for split?
            if random.randint(0, 1) == 0:
                dir_x = random.randint(0, 3)
                dir_y = random.randint(0, 3)
                print('retry split!')
                return
            else:
                cords = np.argwhere(self.canvas > 0)
                self.canvas[cords[:, 0], cords[:, 1]] = 1
                print('no split again!')
                return

        # clockwise order from NW
        if (dir_y == 1) or (dir_y == 3):
            pass
        else:
            if random.randint(0, 1) == 0:
                self.replaceSection(self.canvas, 2, 1)
                self.replaceSection(self.canvas, 3, 1)
            else:
                self.replaceSection(self.canvas, 2, 2 + random.choice([1, -1]))

        if (dir_x == 2) or (dir_x == 3):
            pass
        else:
            if random.randint(0, 1) == 0:
                self.replaceSection(self.canvas, 4, 3)
                self.replaceSection(self.canvas, 5, 3)
            else:
                self.replaceSection(self.canvas, 4, 4 + random.choice([1, -1]))

        if (dir_y == 2) or (dir_y == 3):
            pass
        else:
            if random.randint(0, 1) == 0:
                self.replaceSection(self.canvas, 6, 5)
                self.replaceSection(self.canvas, 7, 5)
            else:
                self.replaceSection(self.canvas, 6, 6 + random.choice([1, -1]))

        if (dir_x == 1) or (dir_x == 3):
            pass
        else:
            if random.randint(0, 1) == 0:
                self.replaceSection(self.canvas, 8, 7)
                self.replaceSection(self.canvas, 1, 7)
            else:
                self.replaceSection(self.canvas, 8, random.choice([1, 7]))

        print(self)

    def decideSplits(self):
        #decide to keep split or remove & merge

        if len(self.getSections()) == 1:
            return None
        
        dirs = self.getBarrierNums()

        # get remaining splits
        for i in dirs:
            if random.randint(0, 1) == 0:  # remove split
                if i == 8:
                    self.replaceSection(self.canvas, i, random.choice([1, 7]))
                else:
                    #self.replaceSection(self.canvas, i, i + random.choice([1, -1]))
                    self.replaceSection(self.canvas, i, i + 1)
                dirs.remove(i)

            else:  # keep split
                pass
        
        ######
        # DECIDE ONE SPLIT

        print(dirs)
        print(self)

    def decideBarriers(self):
        #decide whether to keep remaining split as barrier

        if random.randint(0, 2) != 0:
            # replace all barriers (evens) with 2
            for i in range(2, 10, 2):
                self.replaceSection(self.canvas, i, 2)
                
        if 9 in self.getSections():
            print(self.getBarrierNums())
            #print("found a 9")
            self.replaceSection(self.canvas, 9, self.getBarrierNums()[-1])
        
        print(self)

    def buildFarm(self):
        self.splitGen()
        self.decideSplits()
        self.decideBarriers()
        self.starter = self.canvas

        pad_x = random.randint(
            self.edge_buffer, self.cols - self.edge_buffer - self.canvas.ndim)
        pad_y = random.randint(
            self.edge_buffer, self.rows - self.edge_buffer - self.canvas.ndim)

        self.canvas = np.lib.pad(self.canvas, ((
            pad_y, self.rows - pad_y - self.canvas.shape[0]), (pad_x, self.cols - pad_x - self.canvas.shape[1])), 'edge')
        
        self.SECTION_NUMBERS = self.getBiomeNums() + self.getBarrierNums()
        self.BARRIER_NUMBERS = self.getBarrierNums()
        
        print(self)

        print(self.getBarrierNums())
        print(self.getBiomeNums())
        

        self.intArray2stringArray()
        print('\nBIOMES')

        biomes_list = [Biome(self, self.canvas_strings, str(self.SECTION_NUMBERS[i])) for i in range(len(self.SECTION_NUMBERS))]
        
        print(self.output_emojis())
        
    def intArray2stringArray(self):
        self.canvas_strings = np.array(self.canvas, dtype='U64')

    def output_emojis(self):
        return '\n'.join([emoji.emojize(''.join([str(elem) for elem in row])) for row in self.canvas_strings])

if __name__ == '__main__':

    moji = EmojiGrid()
    moji.buildFarm()

