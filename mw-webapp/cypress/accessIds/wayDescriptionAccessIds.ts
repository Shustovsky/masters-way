export const wayDescriptionAccessIds = {
    wayDashBoardLeft: {
        title : 'wayTitle'
    },

    peopleBlock: {
        applyAsMentorButton: 'applyAsMentorButton',
        mentorOfWayText: 'mentorOfWayText',
        wayMentorLink: 'wayMentorLink',
        childWaysTitle: "childWaysTitle",
        childLink: (name: string) => `childLink_${name}`,
        deleteFromCompositeWayButton: (name: string) => `deleteFromCompositeWayButton_${name}`,
        dialogContent: {
            deleteButton: 'deleteButton'
        }
    },

    mentorRequestDialog: {
        mentorNameLink: 'mentorName',
        acceptButton: 'acceptButton'
    },

    wayActionMenu: {
        wayActionButton: 'wayActionButton',
        wayActionMenuList: "wayActionMenuList",
        wayMenuItem: 'wayMenuItem',
        waySubMenuItem: 'waySubMenuItem',
        waySubTriggerItem: 'waySubTriggerItem'
    }
};