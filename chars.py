#!/usr/local/bin/python3

# MIT License

# Copyright (c) 2017 Clayton Dorsey
# https://cydy.io @claytondorsey 

ANIMALS_FARM = [
    ':cow:',
    ':ox:',
    ':water_buffalo:',
    ':horse:',
    ':pig:',
    ':goat:',
    ':sheep:',
    ':ram:',
    ':rabbit:'
]

ANIMALS_BUGS = [
    ':lady_beetle:',
    ':honeybee:',
    ':butterfly:'
]

ANIMALS_BIRDS = [
    ':front-facing_baby_chick:',
    ':rooster:',
    ':hatching_chick:',
    ':egg:',
    ':duck:',
    ':turkey:'
]

ANIMALS_NON_FARM = [
    ':dog:',
    ':cat:',
    ':owl:',
    ':bat:',
    ':eagle:',
    ':chipmunk:',
    ':rat:',
    ':mouse:',
    ':rabbit:'
]

PLANTS_FIELD = [
    ':seedling:',
    ':herb:',
    ':maple_leaf:',
    ':leaf_fluttering_in_wind:',
    ':fallen_leaf:',
    ':mushroom:',
    ':sheaf_of_rice:',
    ':four_leaf_clover:',
    ':shamrock:'
]

PLANTS_FLOWERS = [
    ':tulip:',
    ':rose:',
    ':rosette:',
    ':blossom:',
    ':wilted_flower:',
    ':sunflower:',
    ':white_flower:',
    ':hibiscus:',
    ':cherry_blossom:'
]

PLANTS_TREES = [
    ':deciduous_tree:',
    ':evergreen_tree:',
    ':palm_tree:'
]

PLANTS_CROPS = [
    ':green_apple:',
    ':red_apple:',
    ':tangerine:',
    ':lemon:',
    ':grapes:',
    ':watermelon:',
    ':strawberry:',
    ':cherries:',
    ':peach:',
    ':tomato:',
    ':eggplant:',
    ':carrot:',
    ':ear_of_corn:',
    ':hot_pepper:',
    ':potato:',
    ':melon:',
    ':pear:',
    ':kiwi_fruit:'
]

BUILDINGS = [
    ':house:',
    ':house_with_garden:',
    ':derelict_house:',
    #':house_buildings:',
    ':Japanese_castle:',
    ':castle:',
    ':tractor:',
    ':delivery_truck:'
]

EARTH = [
    ':mountain:',
    ':snow-capped_mountain:'
]

SKY = [
    ':helicopter:',
    ':small_airplane:',
    ':rainbow:',
    ':high_voltage:',
    ':cloud:',
    ':cloud_with_lightning:',
    ':cloud_with_lightning_and_rain:',
    ':cloud_with_rain:',
    ':cloud_with_snow:'
]

SKY_SUNS = [
    ':sun:',
    ':sun_behind_cloud:',
    ':sun_behind_large_cloud:',
    ':sun_behind_rain_cloud:',
    ':sun_behind_small_cloud:',
    ':sun_with_face:'
]

BARRIERS = [
    ':chains:'
]

BARRIERS = BARRIERS + EARTH + PLANTS_TREES #+ PLANTS_FLOWERS

WATER = [
    ':water_wave:'
]

POND = [
    ':crocodile:',
    ':fish:',
    ':turtle:',
    ':duck:',
    ':spiral_shell:',
    ':shrimp:',
    ':canoe:'
]

PEOPLE = [
    ':cowboy_hat_face:',
    ':robot_face:',
    ':man_farmer:',
    ':man_farmer_dark_skin_tone:',
    ':man_farmer_light_skin_tone:',
    ':man_farmer_medium-dark_skin_tone:',
    ':man_farmer_medium-light_skin_tone:',
    ':man_farmer_medium_skin_tone:',
    ':woman_farmer:',
    ':woman_farmer_dark_skin_tone:',
    ':woman_farmer_light_skin_tone:',
    ':woman_farmer_medium-dark_skin_tone:',
    ':woman_farmer_medium-light_skin_tone:',
    ':woman_farmer_medium_skin_tone:'
]

__ALL__ = ANIMALS_FARM + ANIMALS_BUGS + ANIMALS_BIRDS + ANIMALS_NON_FARM + PLANTS_FIELD + PLANTS_FLOWERS + PLANTS_TREES + PLANTS_CROPS + BUILDINGS + EARTH + SKY + SKY_SUNS + BARRIERS + WATER + POND + PEOPLE

#print(emoji.emojize(''.join(__ALL__)))
