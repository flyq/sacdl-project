function haveBreakfast(food, drink, callback) {
    console.log('Having breakfast of ' + food + ', ' + drink);
    callback();
}

function nextwork() {
    console.log('Finished breakfast. Time to go to work!');
}

haveBreakfast('toast', 'coffee', nextwork)

